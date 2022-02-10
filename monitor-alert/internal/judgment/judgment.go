package judgment

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/proto/judgpb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/pkg/setting"
	"google.golang.org/grpc"
)

type IJudgment interface {
	// Check 对接入服务发来的告警信息进行处理
	Check(agent *model.AgentReport) error
	// RegisterService 注册rpc服务
	RegisterService(ser *grpc.Server)
}

type server struct {
	proxy *service.Service
}

func (s *server) Check(agent *model.AgentReport) error {
	return s.proxy.Check(agent)
}

func (s *server) RegisterService(ser *grpc.Server) {
	judgpb.RegisterRuleUpdaterServer(ser, s.proxy)
}

func NewService(s *setting.Setting, send send.ISend) *server {
	return &server{
		proxy: service.NewService(s, send),
	}
}
