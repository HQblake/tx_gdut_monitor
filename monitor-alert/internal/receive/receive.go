package receive

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/receive/receivepb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/receive/service"
	"google.golang.org/grpc"
)

type IReceive interface {
	// RegisterService 注册rpc服务
	RegisterService(ser *grpc.Server)
}

type server struct {
	proxy *service.ReceiveService
}

func (s *server) RegisterService(ser *grpc.Server) {
	receivepb.RegisterReportServerServer(ser, s.proxy)
}

func NewService(judgment judgment.IJudgment) *server {
	return &server{
		proxy: service.NewService(judgment),
	}
}
