package service

import (
	managepb2 "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/judgment/gen"
	sendpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/send/gen"
	managepb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/store/gen"
	judgpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/service/gen"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/agent"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/http"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/send"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/show"
	grpc "google.golang.org/grpc"
	"log"
	"net"
)

type Admin struct {
	service *service.Service
	ApiHandler *http.Handler
}

func Register(alertConn, storeConn *grpc.ClientConn) *Admin {
	judgmentService := judgment.NewService(managepb.NewJudgmentServiceClient(storeConn), managepb2.NewRuleUpdaterClient(alertConn))
	sendService := send.NewService(sendpb.NewSendServiceClient(alertConn), managepb.NewSendServiceClient(storeConn))
	showService := show.NewService(managepb.NewHistoryServiceClient(storeConn), managepb.NewMetricServiceClient(storeConn))
	agentService := agent.NewService(managepb.NewAgentServiceClient(storeConn), sendService)
	a :=  &Admin{
		service: service.NewService(judgmentService),
		ApiHandler: http.NewHandler(agentService, judgmentService, sendService, showService),
	}

	// 发送服务init
	sendService.Init()

	return a
}

func (a *Admin) RegisterService(lis net.Listener){
	go func() {
		ser := grpc.NewServer()
		// 管理服务开启grpc通信
		judgpb.RegisterManageServiceServer(ser, a.service)
		err := ser.Serve(lis)
		defer ser.Stop()
		if err != nil {
			log.Fatalf("manage grpc server error %s", err.Error())
		}
	}()
}
