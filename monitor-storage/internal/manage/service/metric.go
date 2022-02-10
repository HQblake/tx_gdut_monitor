package service

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/dao"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/manage/managepb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
)

type MetricService struct {
	dao *dao.StorageDao
	*managepb.UnimplementedMetricServiceServer
}

func (m *MetricService) GetMetricData(request *managepb.MetricRequest, server managepb.MetricService_GetMetricDataServer) error {
	metrics, err := m.dao.GetMetricData(request.IP, request.Local, request.Metric,
		request.Period, request.Begin, request.End, request.Method, request.Limit)
	if err != nil {
		_ = server.Send(&managepb.MetricResponse{
			Code: managepb.ResponseCode_ERROR,
			Msg:  err.Error(),
		})
		return err
	}

	for _, metric := range metrics {
		_ = server.Send(&managepb.MetricResponse{
			Code: managepb.ResponseCode_SUCCESS,
			Msg:  "SUCCESS",
			Result: &managepb.MetricResult{
				Timestamp: metric.Timestamp,
				Metric:    metric.Name,
				Value:     metric.Value,
			},
		})
	}
	return nil
}

func NewMetricService(s *setting.Setting) *MetricService {
	return &MetricService{
		dao: dao.NewStorageDao(s),
	}
}
