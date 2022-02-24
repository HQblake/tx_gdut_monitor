package roll

import (
	"fmt"
	model2 "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"log"
	"sync"
	"time"
)

// 滚动告警收敛算法，尚未完善
const (
	convergenceMinInterval int64 = 60 * 5
	convergenceMapLimit    int   = 300
	convergenceMaxInterval int64 = 60 * 40

)

type convergenceData struct {
	scrollStartTime         int64
	scrollLeft, scrollRight int64
}
type Convergence struct {
	stack    *model.Queue
	manage   output.IManager
	lock     sync.RWMutex
	data     map[string]*convergenceData
	timer    *time.Ticker
	infoPool *sync.Pool
}

func NewConvergence(manage output.IManager, interval time.Duration) *Convergence {
	c := &Convergence{
		stack:  model.InitQueue(),
		manage: manage,
		lock:   sync.RWMutex{},
		data:   make(map[string]*convergenceData),
		infoPool: &sync.Pool{
			New: func() interface{} {
				return model.Info{}
			},
		},
		timer: time.NewTicker(interval),
	}
	go c.SendQueueTask()
	return c
}

func (c *Convergence) Alert(alert *model2.AlertInfo) error{
	target := fmt.Sprintf("%s-%s", alert.IP, alert.Local)
	for metric, info := range alert.Metrics {
		i := c.newInfo(target, info)
		go c.alertConvergence(target, i, time.Now(), metric)
	}
	return nil

}

func (c *Convergence) alertConvergence(target string, info model.Info, timestamp time.Time, metric string) {
	// 滚动收敛算法，暂时先实现简单的收敛处理
	c.lock.Lock()
	if _, ok := c.data[metric]; !ok {
		c.data[metric] = &convergenceData{scrollStartTime: timestamp.Unix(), scrollLeft: 0, scrollRight: 1}
		if len(c.data) > convergenceMapLimit {
			c.clearConvergenceMap()
		}
	}
	cod := c.data[metric]

	//convergence and update scroll if need
	duration := timestamp.Unix() - cod.scrollStartTime
	repeatAlter := duration < cod.scrollLeft*convergenceMinInterval
	if repeatAlter {
		c.lock.Unlock()
		return
	}
	cod.scroll(timestamp)
	c.lock.Unlock()
	c.lock.Lock()
	data := model.Node{
		Target:  target,
		Message: info,
		Metric:  metric,
	}
	c.push(data)
	c.lock.Unlock()
}

func (c *Convergence) push(data model.Node) {
	c.stack.EnQueue(data)
}

func (c *Convergence) pop() model.Node {
	return c.stack.DeQueue()
}

func (cod *convergenceData) scroll(timestamp time.Time) {
	reset := (timestamp.Unix() - cod.scrollStartTime) > cod.scrollRight*convergenceMinInterval
	if reset {
		cod.scrollLeft, cod.scrollRight, cod.scrollStartTime = 0, 1, timestamp.Unix()
		return
	}

	reachMax := cod.scrollRight*convergenceMinInterval >= convergenceMaxInterval
	if reachMax {
		cod.scrollLeft, cod.scrollRight = cod.scrollRight, cod.scrollRight+(cod.scrollRight-cod.scrollLeft)
	} else {
		cod.scrollLeft, cod.scrollRight = cod.scrollRight, cod.scrollRight+(cod.scrollRight-cod.scrollLeft)*2
	}

}

func (c *Convergence) clearConvergenceMap() {
	haveExpired := false
	var minStartTime int64
	minStartTimeKey := ""
	expires := make([]string, 0)
	for k, v := range c.data {
		// get expired array and oldest time
		if time.Now().Unix()-v.scrollStartTime > v.scrollLeft*convergenceMinInterval {
			haveExpired = true
			expires = append(expires, k)
		}
		if 0 == minStartTime || v.scrollStartTime < minStartTime {
			minStartTime = v.scrollStartTime
			minStartTimeKey = k
		}
	}
	// delete expired or oldest time if need
	if haveExpired {
		for _, k := range expires {
			delete(c.data, k)
		}
	} else {
		delete(c.data, minStartTimeKey)
	}
}

func (c *Convergence) SendQueueTask() {
	go func() {
		for {
			select {
			case <-c.timer.C:
				aggregation := c.merge()
				if aggregation == nil {
					continue
				}
				for k, v := range aggregation {
					c.send(k, v)
				}

			}
		}
	}()
}

// 获取所有的告警信息
func (c *Convergence) allData() []model.Node {
	list := make([]model.Node, 0)
	c.lock.Lock()
	for {
		data := c.pop()
		if data.Target == "" {
			break
		}
		list = append(list, data)
	}
	c.lock.Unlock()
	return list
}

// 聚合告警消息
func (c *Convergence) merge() map[string][]model.Info {
	all := c.allData()
	if all == nil {
		return nil
	}
	arrange := make(map[string][]model.Info)
	for _, node := range all {
		if value, ok := arrange[node.Target]; ok {
			arrange[node.Target] = append(value, node.Message)
		} else {
			r := make([]model.Info, 0, 10)
			r = append(r, node.Message)
			arrange[node.Target] = r
		}
	}
	return arrange
}

// 发送处理
func (c *Convergence) send(agent string, infos []model.Info) {
	outputs := c.manage.GetOutputs(agent)
	err := outputs.Output(infos)
	if err != nil {
		log.Println(err)
	}
	c.releaseAll(infos)
}

func (c *Convergence) newInfo(agent string, alert model2.MetricInfo) model.Info {
	i := c.infoPool.Get().(model.Info)
	i.Agent = agent
	i.Metric = alert.Metric
	i.Value = alert.Value
	i.Threshold = alert.Threshold
	i.Level = output.Level(alert.Level).String()
	i.Duration = alert.Duration
	i.ParseDate(alert.Start)
	i.ParseMethod(alert.Method)
	return i
}

func (c *Convergence) releaseAll(infos []model.Info) {
	for _, info := range infos {
		c.release(info)
	}
}

func (c *Convergence) release(info model.Info) {
	c.infoPool.Put(info)
}
