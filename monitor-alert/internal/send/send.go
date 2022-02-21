package send

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	sendpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/api/gen"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/api/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/convergence"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/convergence/simple"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"google.golang.org/grpc"
	"sync"
	"time"
)

// ISend 接入判定服务
type ISend interface {
	// Send 对判定服务发来的告警信息进行处理，进行通知
	Send(alert *model.AlertInfo) error
	// RegisterService 注册rpc服务
	RegisterService(ser *grpc.Server)
}

type Service struct {
	proxy    *service.Service
	convergence convergence.IConvergence
	infoPool *sync.Pool
}



// NewService 初始化发送服务，提供对外的判定服务结构（判定服务直接调用该结构的Send方法即可完成发送）
func NewService() *Service{
	Register()
	agents := output.NewManager()
	return &Service{
		// 简单收敛聚合处理，一分钟告警一次
		convergence: simple.NewConvergence(agents, time.Minute * 1),
		// 滚动收敛聚合处理，一分钟告警一次
		//convergence: roll.NewConvergence(agents, time.Minute * 1),
		proxy:  service.NewService(agents),
	}
}
func (s *Service) RegisterService(ser *grpc.Server) {
	sendpb.RegisterSendServiceServer(ser, s.proxy)
}

func (s *Service) Send(alert *model.AlertInfo) error {
	return s.convergence.Alert(alert)
}

