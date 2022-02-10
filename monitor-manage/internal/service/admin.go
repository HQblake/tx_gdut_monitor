package service

import (
	judgpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/service/gen"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/agent"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/http"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/send"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/show"
)

type Admin struct {
	Service *service.Service
	ApiHandler *http.Handler
}

func Register() *Admin {
	agent := agent.NewService()
	judgment := judgment.NewService()
	send := send.NewService()
	send.Init()
	show := show.NewService()
	// 发送服务init
	// 管理服务开启grpc通信
	a :=  &Admin{
		Service: service.NewService(judgment),
		ApiHandler: http.NewHandler(agent, judgment, send, show),
	}
	judgpb.RegisterManageServiceServer(a.Service)
	return a
}
