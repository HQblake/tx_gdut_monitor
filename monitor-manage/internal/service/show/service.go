package show

import (
	"context"
	"io"
	"log"

	// "encoding/json"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"
	managepb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/store/gen"
)

// Service 实现IShow接口的实例，包括数据处理，最后调用存储模块的rpc服务
type Service struct {
	// 获取历史告警信息的客户端
	alertClient managepb.HistoryServiceClient
	// 获取实时监控指标信息的客户端
	metricClient managepb.MetricServiceClient
}

func NewService(alertClient managepb.HistoryServiceClient, metricClient managepb.MetricServiceClient) *Service {
	return &Service{
		alertClient:  alertClient,
		metricClient: metricClient,
	}
}

func (s *Service) GetWarnInfo() ([]model.HistoryInfo, error) {
	var err error
	// 获取告警服务中所有的告警信息
	stream, err := s.alertClient.GetAllAlertInfo(context.Background(), &managepb.BaseRequest{})
	if err != nil {
		return nil, err
	}
	var resp *managepb.AlertResponse
	res := make([]model.HistoryInfo, 0, 10)
	// 遍历获取指定metric的规则，没有则用默认规则代替
	for {
		resp, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("grpc get history rule error %v", err)
			continue
		}
		if resp.Code != managepb.ResponseCode_SUCCESS {
			log.Printf("grpc get history rule error %v", resp.Msg)
			continue
		}
		config := resp.GetResult()
		conf := model.HistoryInfo{
			Id:        config.GetID(),
			Ip:        config.GetIP(),
			Local:     config.GetLocal(),
			Metric:    config.GetMetric(),
			Value:     config.Value(),
			Method:    config.GetMethod(),
			Level:     config.GetLevel(),
			Threshold: config.GetThreshold(),
			Start:     config.GetStart(),
			Duration:  config.GetDuration(),
		}
		res = append(res, conf)
	}
	return res, nil
}

func (s *Service) GetMetricsInOneDay(ip string, local string, metric string) []model.MetricsInfo {
	panic("implement me")
}
