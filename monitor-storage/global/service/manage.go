package service

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/manage/managepb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/manage/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
	"google.golang.org/grpc"
)

type ManageService struct {
	agentProxy    *service.AgentService
	historyProxy  *service.HistoryService
	judgmentProxy *service.JudgmentService
	metricProxy   *service.MetricService
	sendProxy     *service.SendService
}

func NewManageService(s *setting.Setting) *ManageService {
	return &ManageService{
		agentProxy:    service.NewAgentService(s),
		historyProxy:  service.NewHistoryService(s),
		judgmentProxy: service.NewJudgmentService(s),
		metricProxy:   service.NewMetricService(s),
		sendProxy:     service.NewSendService(s),
	}
}

func (ms *ManageService) RegisterManageService(server *grpc.Server) {
	managepb.RegisterAgentServiceServer(server, ms.agentProxy)
	managepb.RegisterHistoryServiceServer(server, ms.historyProxy)
	managepb.RegisterJudgmentServiceServer(server, ms.judgmentProxy)
	managepb.RegisterMetricServiceServer(server, ms.metricProxy)
	managepb.RegisterSendServiceServer(server, ms.sendProxy)
}
