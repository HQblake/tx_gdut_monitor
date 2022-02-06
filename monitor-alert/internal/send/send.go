package send

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/api/service"
	model2 "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"log"
	"sync"
)

// ISend 接入判定服务
type ISend interface {
	// Send 对判定服务发来的告警信息进行处理，进行通知
	Send(alert model.AlertInfo) error
}

type Service struct {
	proxy *service.Service
	agents output.IManager
	infoPool *sync.Pool
}

// NewService 初始化发送服务，提供对外的判定服务结构（判定服务直接调用该结构的Send方法即可完成发送）
func NewService(addr string) (*Service, error) {
	Register()
	agents := output.NewManager()
	s, err := service.NewService(agents, addr)
	if err != nil {
		return nil, err
	}
	return &Service{
		agents: agents,
		proxy:  s,
		infoPool: &sync.Pool{
			New: func() interface{} {
				return model2.Info{}
			},
		},
	}, nil
}

func (s *Service) Send(alert model.AlertInfo) error {
	outputs := s.agents.GetOutputs(alert.AgentID)
	for _, info := range alert.Metrics {
		i := s.newInfo(alert.AgentID, info)
		err := outputs.Output(i)
		if err != nil {
			log.Println(err)
		}
		s.release(i)
	}
	return nil
}

func (s *Service) newInfo(agent string, alert model.MetricInfo) model2.Info {
	i := s.infoPool.Get().(model2.Info)
	i.Agent = agent
	i.Metric = alert.Metric
	i.Value = alert.Value
	i.Threshold = alert.Threshold
	i.Level = output.Level(alert.Level).String()
	i.Duration = alert.Duration
	i.Start = alert.Start
	i.ParseMethod(alert.Method)
	return i
}

func (s *Service) release(info model2.Info) {
	s.infoPool.Put(info)
}


