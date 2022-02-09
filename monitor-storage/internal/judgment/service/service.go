package service

import (
	"context"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/dao"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/judgment/judgpb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
)

type MetricService struct {
	dao *dao.StorageDao
	*judgpb.UnimplementedMetricServiceServer
}

func (m *MetricService) GetAggregatedData(ctx context.Context, request *judgpb.AggregatedRequest) (*judgpb.AggregatedResponse, error) {
	metric := model.MetricPool.Get().(*model.Metric)
	defer model.MetricPool.Put(metric)
	parseMetric(request, metric)
	// 1. 执行事务处理，获取聚合结果
	value, err := m.dao.GetAggregatedData(request.Period, request.Method, metric)
	if err != nil {
		return &judgpb.AggregatedResponse{Code: judgpb.ResponseCode_ERROR, Msg: err.Error()}, err
	}

	// 2. 在MySQL中进行相关记录
	err = m.dao.SaveAgentInfo(metric)
	if err != nil {
		return &judgpb.AggregatedResponse{Code: judgpb.ResponseCode_ERROR, Msg: err.Error()}, err
	}
	return &judgpb.AggregatedResponse{Code: judgpb.ResponseCode_SUCCESS, Msg: "SUCCESS", Result: value}, nil
}

func (m *MetricService) InsertAlertInfo(ctx context.Context, request *judgpb.HistoryInfoRequest) (*judgpb.HistoryInfoResponse, error) {
	history := model.HistoryInfoPool.Get().(*model.HistoryInfo)
	defer model.HistoryInfoPool.Put(history)
	parseHistory(request, history)
	err := m.dao.SaveAlertInfo(history)
	if err != nil {
		return &judgpb.HistoryInfoResponse{Code: judgpb.ResponseCode_ERROR, Msg: err.Error()}, err
	}
	return &judgpb.HistoryInfoResponse{Code: judgpb.ResponseCode_SUCCESS, Msg: "SUCCESS"}, nil
}

func NewService(s *setting.Setting) *MetricService {
	return &MetricService{
		dao: dao.NewStorageDao(s),
	}
}

func parseMetric(request *judgpb.AggregatedRequest, metric *model.Metric) {
	metric.Name = request.Metric
	metric.Value = request.Value
	metric.IP = request.IP
	metric.Local = request.Local
	metric.Port = request.Port
	metric.Timestamp = request.Timestamp
	metric.Dimensions = request.Dimensions
}

func parseHistory(request *judgpb.HistoryInfoRequest, history *model.HistoryInfo) {
	history.IP = request.IP
	history.Local = request.Local
	history.Metric = request.Metric
	history.Value = request.Value
	history.Threshold = request.Threshold
	history.Method = request.Method
	history.Level = request.Level
	history.Start = request.Start
	history.Duration = request.Duration
}
