package service

import (
	"encoding/json"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/dao"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/proto"
	"golang.org/x/net/context"
)

type ManageService struct{}

func (m *ManageService) GetAllAgentInfo(request *proto.AgentRequest, server proto.ManageConfig_GetAllAgentInfoServer) error {
	agents := dao.GetAllAgentInfo()
	for _, agent := range agents {
		result, _ := json.Marshal(agent)
		_ = server.Send(&proto.AgentResponse{Code: 200, Msg: "SUCCESS", Result: string(result)})
	}
	return nil
}

func (m *ManageService) GetMetricsByAgentID(request *proto.AgentRequest, server proto.ManageConfig_GetMetricsByAgentIDServer) error {
	metrics := dao.GetMetricsByAgentID(request.AgentID)
	for _, metric := range metrics {
		_ = server.Send(&proto.AgentResponse{Code: 200, Msg: "SUCCESS", Result: metric})
	}
	return nil
}

func (m *ManageService) GetMetricData(request *proto.MetricRequest, server proto.ManageConfig_GetMetricDataServer) error {
	metrics := dao.GetMetricData(request.AgentID, request.Metric, request.Period, request.Begin, request.End)
	for _, metric := range metrics {
		_ = server.Send(&proto.MetricResponse{
			Code:      200,
			Msg:       "SUCCESS",
			Timestamp: metric.Timestamp,
			Metric:    metric.Name,
			Value:     float32(metric.Value),
		})
	}
	return nil
}

func (m *ManageService) GetAllAlertInfo(request *proto.AlertRequest, server proto.ManageConfig_GetAllAlertInfoServer) error {
	alerts := dao.GetAllAlertInfo()
	for _, alert := range alerts {
		for _, info := range alert.Metrics {
			_ = server.Send(&proto.AlertResponse{
				Code:      200,
				Msg:       "SUCCESS",
				ID:        info.ID,
				AgentID:   alert.AgentID,
				Metric:    info.Metric,
				Value:     float32(info.Value),
				Threshold: float32(info.Threshold),
				Duration:  info.Duration,
				Level:     int32(info.Level),
				Begin:     info.Start,
			})
		}
	}
	return nil
}

func (m *ManageService) GetAlertInfo(request *proto.AlertRequest, server proto.ManageConfig_GetAlertInfoServer) error {
	alerts := dao.GetAlertInfo(request.ID, request.Level, request.AgentID, request.Metric, request.Begin, request.End)
	for _, alert := range alerts {
		for _, info := range alert.Metrics {
			_ = server.Send(&proto.AlertResponse{
				Code:      200,
				Msg:       "SUCCESS",
				ID:        info.ID,
				AgentID:   alert.AgentID,
				Metric:    info.Metric,
				Value:     float32(info.Value),
				Threshold: float32(info.Threshold),
				Duration:  info.Duration,
				Level:     int32(info.Level),
				Begin:     info.Start,
			})
		}
	}
	return nil
}

func (m *ManageService) DelAlterInfo(ctx context.Context, request *proto.AlertRequest) (*proto.AlertResponse, error) {
	err := dao.DelAlterInfo(request.ID)
	if err != nil {
		return &proto.AlertResponse{Code: 500, Msg: err.Error()}, err
	}
	return &proto.AlertResponse{Code: 200, Msg: "SUCCESS"}, nil
}

func (m *ManageService) AddConfig(ctx context.Context, request *proto.ConfigRequest) (*proto.ConfigResponse, error) {
	if request.Type == proto.ConfigType_ALERT {

	} else {

	}

	return &proto.ConfigResponse{Code: 200, Msg: "SUCCESS"}, nil
}

func (m *ManageService) UpdateConfigById(ctx context.Context, request *proto.ConfigRequest) (*proto.ConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ManageService) DeleteConfigById(ctx context.Context, request *proto.ConfigRequest) (*proto.ConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ManageService) GetConfig(request *proto.ConfigRequest, server proto.ManageConfig_GetConfigServer) error {
	//TODO implement me
	panic("implement me")
}

func (m *ManageService) GetAllConfigs(request *proto.ConfigRequest, server proto.ManageConfig_GetAllConfigsServer) error {
	//TODO implement me
	panic("implement me")
}

func NewManageService() (*ManageService, error) {
	return &ManageService{}, nil
}
