package service

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/dao"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/proto"
	"golang.org/x/net/context"
)

type MetricService struct{}

func (m *MetricService) InsertAlertInfo(ctx context.Context, request *proto.AlertInfoRequest) (*proto.AlertInfoResponse, error) {
	alert := model.AlertInfoPool.Get().(*model.AlertInfo)
	defer model.AlertInfoPool.Put(alert)
	parseAlertInfo(request, alert)
	err := dao.SaveAlertInfo(alert)
	if err != nil {
		return &proto.AlertInfoResponse{Code: 500, Msg: err.Error()}, err
	}
	return &proto.AlertInfoResponse{Code: 200, Msg: "SUCCESS"}, nil
}

func (m *MetricService) GetAggregatedData(ctx context.Context, request *proto.AggregatedRequest) (*proto.AggregatedResponse, error) {
	metric := model.MetricPool.Get().(*model.Metric)
	defer model.MetricPool.Put(metric)
	parseMetric(request, metric)
	// 1. 执行事务处理，获取聚合结果
	value, err := dao.GetAggregatedData(request.Period, request.Method, metric)
	if err != nil {
		return &proto.AggregatedResponse{Code: 500, Msg: err.Error()}, err
	}

	// 2. 在MySQL中进行相关记录
	err = dao.SetAgentInfo(metric)
	if err != nil {
		return &proto.AggregatedResponse{Code: 500, Msg: err.Error()}, err
	}
	return &proto.AggregatedResponse{Code: 200, Msg: "SUCCESS", Value: value}, nil
}

func parseAlertInfo(request *proto.AlertInfoRequest, alert *model.AlertInfo) {
	alert.AgentID = request.AgenID
	alert.Metrics = make(map[string]model.MetricInfo)
	for k, v := range request.Metrics {
		alert.Metrics[k] = model.MetricInfo{
			Metric:    v.Metric,
			Value:     v.Value,
			Threshold: v.Threshold,
			Method:    int8(v.Method),
			Level:     int8(v.Level),
			Duration:  v.Duration,
			Start:     v.Start,
		}
	}
}

func parseMetric(request *proto.AggregatedRequest, metric *model.Metric) {
	metric.Name = request.Metric
	metric.Value = request.Value
	metric.IP = request.IP
	metric.Local = request.Local
	metric.Port = request.Port
	metric.Timestamp = request.Timestamp
	metric.Dimensions = request.Dimensions
}

func NewMetricService() (*MetricService, error) {
	return &MetricService{}, nil
}
