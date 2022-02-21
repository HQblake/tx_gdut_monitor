package aggregation

// 告警数据聚合操作，避免同一时间段频繁告警

import (
	"fmt"
	model2 "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"log"
	"sync"
	"time"
)


type Convergence struct {
	stack    *model.Queue
	manage   output.IManager
	lock     sync.RWMutex
	timer    *time.Ticker
	infoPool *sync.Pool
}

// NewConvergence interval 报警频率
func NewConvergence(manage output.IManager, interval time.Duration) *Convergence {
	c := &Convergence{
		stack:  model.InitQueue(),
		manage: manage,
		lock:   sync.RWMutex{},
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

func (c *Convergence) Alert(alert *model2.AlertInfo) error {
	target := fmt.Sprintf("%s-%s", alert.IP, alert.Local)
	c.lock.Lock()
	for metric, info := range alert.Metrics {
		data := model.Node{
			Target:  target,
			Message: c.newInfo(target, info),
			Metric:  metric,
		}
		c.push(data)
	}
	c.lock.Unlock()
	return nil
}


func (c *Convergence) push(data model.Node) {
	c.stack.EnQueue(data)
}

func (c *Convergence) pop() model.Node {
	return c.stack.DeQueue()
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
	i.Start = time.Unix(alert.Start, 0).Format("[2006-01-01 15:04:05]")
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

