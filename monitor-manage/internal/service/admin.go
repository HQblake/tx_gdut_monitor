package service

import (
	managepb2 "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/judgment/gen"
	sendpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/send/gen"
	managepb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/store/gen"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/agent"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/http"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/send"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/show"
	grpc "google.golang.org/grpc"
)

type Admin struct {
	Service *service.Service
	ApiHandler *http.Handler
}

func Register(alertConn, storeConn *grpc.ClientConn) *Admin {
	agentService := agent.NewService(managepb.NewAgentServiceClient(storeConn))
	judgmentService := judgment.NewService(managepb.NewJudgmentServiceClient(storeConn), managepb2.NewRuleUpdaterClient(alertConn))
	sendService := send.NewService(sendpb.NewSendServiceClient(alertConn), managepb.NewSendServiceClient(storeConn))
	showService := show.NewService(managepb.NewHistoryServiceClient(storeConn), managepb.NewMetricServiceClient(storeConn))
	a :=  &Admin{
		Service: service.NewService(judgmentService),
		ApiHandler: http.NewHandler(agentService, judgmentService, sendService, showService),
	}

	// 发送服务init
	sendService.Init()

	return a
}
