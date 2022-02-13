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

// 获取截止目前为止所有的告警信息
func (s *Service) GetWarnInfo() ([]model.HistoryInfo, error) {
	var err error
	stream, err := s.alertClient.GetAllAlertInfo(context.Background(), &managepb.BaseRequest{})
	if err != nil {
		return nil, err
	}
	var resp *managepb.AlertResponse
	res := make([]model.HistoryInfo, 0, 10)
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
			Value:     config.GetValue(),
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

// 根据告警等级查看当天的所有告警信息
func (s *Service) GetWarnInfoWithLevel(level int32) ([]model.HistoryInfo, error) {
	var err error
	stream, err := s.alertClient.GetAlertInfo(context.Background(), &managepb.AlertRequest{Level: level})
	if err != nil {
		return nil, err
	}
	var resp *managepb.AlertResponse
	res := make([]model.HistoryInfo, 0, 10)
	for {
		resp, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("grpc get historyAlert rule error %v", err)
			continue
		}
		if resp.Code != managepb.ResponseCode_SUCCESS {
			log.Printf("grpc get historyAlert rule error %v", resp.Msg)
			continue
		}
		config := resp.GetResult()
		conf := model.HistoryInfo{
			Id:        config.GetID(),
			Ip:        config.GetIP(),
			Local:     config.GetLocal(),
			Metric:    config.GetMetric(),
			Value:     config.GetValue(),
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

// 查看某天的所有告警信息
func (s *Service) GetWarnInfoWithTimestamp(timeStamp int64) ([]model.HistoryInfo, error) {
	var err error
	stream, err := s.alertClient.GetAlertInfo(context.Background(), &managepb.AlertRequest{Begin: timeStamp})
	if err != nil {
		return nil, err
	}
	var resp *managepb.AlertResponse
	res := make([]model.HistoryInfo, 0, 10)
	for {
		resp, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("grpc get historyAlert rule error %v", err)
			continue
		}
		if resp.Code != managepb.ResponseCode_SUCCESS {
			log.Printf("grpc get historyAlert rule error %v", resp.Msg)
			continue
		}
		config := resp.GetResult()
		conf := model.HistoryInfo{
			Id:        config.GetID(),
			Ip:        config.GetIP(),
			Local:     config.GetLocal(),
			Metric:    config.GetMetric(),
			Value:     config.GetValue(),
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

// 根据agentId(ip + local)获取当前主机所有的告警信息
func (s *Service) GetWarnInfoWithId(ip string, local string) ([]model.HistoryInfo, error) {
	var err error
	stream, err := s.alertClient.GetAlertInfo(context.Background(), &managepb.AlertRequest{IP: ip, Local: local})
	if err != nil {
		return nil, err
	}
	var resp *managepb.AlertResponse
	res := make([]model.HistoryInfo, 0, 10)
	for {
		resp, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("grpc get historyAlert rule error %v", err)
			continue
		}
		if resp.Code != managepb.ResponseCode_SUCCESS {
			log.Printf("grpc get historyAlert rule error %v", resp.Msg)
			continue
		}
		config := resp.GetResult()
		conf := model.HistoryInfo{
			Id:        config.GetID(),
			Ip:        config.GetIP(),
			Local:     config.GetLocal(),
			Metric:    config.GetMetric(),
			Value:     config.GetValue(),
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

// 根据开始时间和数据量限制获取指定条数的指标数据
func (s *Service) GetMetricsWithTime(ip string, local string, metric string, begin int64, limit int32) ([]model.MetricsInfo, error) {
	var err error
	stream, err := s.metricClient.GetMetricData(context.Background(), &managepb.MetricRequest{IP: ip, Local: local, Metric: metric, Begin: begin, Limit: limit})
	if err != nil {
		return nil, err
	}
	var resp *managepb.MetricResponse
	res := make([]model.MetricsInfo, 0, 10)
	for {
		resp, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("grpc get historyAlert rule error %v", err)
			continue
		}
		if resp.Code != managepb.ResponseCode_SUCCESS {
			log.Printf("grpc get historyAlert rule error %v", resp.Msg)
			continue
		}
		config := resp.GetResult()
		conf := model.MetricsInfo{
			Timestamp: config.GetTimestamp(),
			Metric:    config.GetMetric(),
			Value:     config.GetValue(),
		}
		res = append(res, conf)
	}
	return res, nil
}
