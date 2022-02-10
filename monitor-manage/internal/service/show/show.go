package show

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"
	managepb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/store/gen"
)

// IShow 图表展示的方案设计，可由自青设计
type IShow interface {
	// GetMetricsInOneDay 比如根据agentId(ip和local)和metric获取一天内的指标情况等
	GetMetricsInOneDay(ip string, local string, metric string) []model.MetricsInfo
}

// Service 实现IShow接口的实例，包括数据处理，最后调用存储模块的rpc服务
type Service struct {
	// 获取历史告警信息的客户端
	alertClient managepb.HistoryServiceClient
	// 获取实时监控指标信息的客户端
	metricClient managepb.MetricServiceClient
}

func (s *Service) GetMetricsInOneDay(ip string, local string, metric string) []model.MetricsInfo {
	panic("implement me")
}

func NewService(alertClient managepb.HistoryServiceClient, metricClient managepb.MetricServiceClient) *Service {
	return &Service{
		alertClient: alertClient,
		metricClient: metricClient,
	}
}
