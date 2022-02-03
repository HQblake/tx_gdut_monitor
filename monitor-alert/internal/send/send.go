package send

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/api/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"log"
)

// ISend 接入判定服务
type ISend interface {
	// Send 对判定服务发来的告警信息进行处理，进行通知
	Send(alert model.AlertInfo) error
}


type Service struct {
	proxy *service.Service
	agents output.IManager
}

func NewService() (*Service, error) {
	InitFactory()
	agents := output.NewManager()
	s, err := service.NewService(agents, ":8082")
	if err != nil {
		return nil, err
	}
	return &Service{
		agents: agents,
		proxy:  s,
	}, nil
}

func (s *Service) Send(alert model.AlertInfo) error {
	outputs := s.agents.GetOutputs(alert.AgentID)
	for _, info := range alert.Metrics {
		err := outputs.Output(info)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
