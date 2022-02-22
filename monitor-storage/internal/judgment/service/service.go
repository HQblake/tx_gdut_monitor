package service

import (
	"context"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/dao"
	judgpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/judgment/judgpb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
	"github.com/panjf2000/ants/v2"
	"log"
	"os"
	"time"
)

type MetricService struct {
	dao  *dao.StorageDao
	pool *ants.Pool
	*judgpb.UnimplementedMetricServiceServer
}

func (m *MetricService) GetAggregatedData(ctx context.Context, request *judgpb.AggregatedRequest) (*judgpb.AggregatedResponse, error) {
	metric := model.MetricPool.Get().(*model.Metric)
	parseMetric(request, metric)
	// 1. 执行事务处理，获取聚合结果
	value, err := m.dao.GetAggregatedData(request.Period, request.Method, metric)
	if err != nil {
		return &judgpb.AggregatedResponse{Code: judgpb.BaseResponseCode_ERRORCODE, Msg: err.Error()}, err
	}

	// 将存储AgentInfo的任务提交至协程池
	err = m.dao.SaveAgentInfo(metric)
	if err != nil {
		log.Printf("Save AgentInfo error: %v\n", err)
	}
	model.MetricPool.Put(metric)

	return &judgpb.AggregatedResponse{Code: judgpb.BaseResponseCode_SUCCESSCODE, Msg: "SUCCESS", Result: value}, nil
}

func (m *MetricService) InsertAlertInfo(ctx context.Context, request *judgpb.HistoryInfoRequest) (*judgpb.HistoryInfoResponse, error) {
	history := model.HistoryInfoPool.Get().(*model.HistoryInfo)
	defer model.HistoryInfoPool.Put(history)
	parseHistory(request, history)
	err := m.dao.SaveAlertInfo(history)
	if err != nil {
		return &judgpb.HistoryInfoResponse{Code: judgpb.BaseResponseCode_ERRORCODE, Msg: err.Error()}, err
	}
	return &judgpb.HistoryInfoResponse{Code: judgpb.BaseResponseCode_SUCCESSCODE, Msg: "SUCCESS"}, nil
}

func NewService(s *setting.Setting) *MetricService {
	ws := &model.WorkersSetting{}
	s.ReadSection("Workers", ws)

	duration, _ := time.ParseDuration(ws.ExpiryDuration)

	// 初始化协程池
	pool, _ := ants.NewPool(ws.Capacity, ants.WithOptions(ants.Options{
		ExpiryDuration:   duration,
		PreAlloc:         ws.PreAlloc,
		MaxBlockingTasks: ws.MaxBlockingTasks,
		Nonblocking:      ws.Nonblocking,
		PanicHandler:     nil,
		Logger:           ants.Logger(log.New(os.Stderr, "", log.LstdFlags)),
	}))
	return &MetricService{
		dao:  dao.NewStorageDao(s),
		pool: pool,
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
