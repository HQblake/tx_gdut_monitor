package instant

// 不做收敛，即时告警

import (
	"fmt"
	model2 "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"log"
	"sync"
)


type Convergence struct {
	manage   output.IManager
	infoPool *sync.Pool
}

// NewConvergence interval 报警频率
func NewConvergence(manage output.IManager) *Convergence {
	c := &Convergence{
		manage: manage,
		infoPool: &sync.Pool{
			New: func() interface{} {
				return model.Info{}
			},
		},
	}
	return c
}

func (c *Convergence) Alert(alert *model2.AlertInfo) error {
	fmt.Printf("开始告警：%+v",alert)
	agent := fmt.Sprintf("%s-%s", alert.IP, alert.Local)
	outputs := c.manage.GetOutputs(agent)
	for _, info := range alert.Metrics {
		i := c.newInfo(agent, info)
		err := outputs.Output([]model.Info{i})
		if err != nil {
			log.Println(err)
		}
		c.release(i)
	}
	return nil
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


func (c *Convergence) release(info model.Info) {
	c.infoPool.Put(info)
}

