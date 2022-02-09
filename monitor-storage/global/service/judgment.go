package service

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/judgment/judgpb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/judgment/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
	"google.golang.org/grpc"
)

type JudgmentService struct {
	proxy *service.MetricService
}

func NewJudgmentService(s *setting.Setting) *JudgmentService {
	return &JudgmentService{
		proxy: service.NewService(s),
	}
}

func (js *JudgmentService) RegisterJudgmentService(server *grpc.Server) {
	judgpb.RegisterMetricServiceServer(server, js.proxy)
}
