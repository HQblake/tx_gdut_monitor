package service

import (
	"context"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/dao"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/manage/managepb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
)

type SendService struct {
	dao *dao.StorageDao
	*managepb.UnimplementedSendServiceServer
}

func (s *SendService) GetAllConfigs(request *managepb.BaseRequest, server managepb.SendService_GetAllConfigsServer) error {
	alerts := s.dao.GetAllAlertConfig()
	for _, config := range alerts {
		_ = server.Send(&managepb.SendConfigResponse{
			Code: managepb.ResponseCode_SUCCESS,
			Msg:  "SUCCESS",
			Result: &managepb.SendEntry{
				ID:       config.ID,
				IP:       config.IP,
				Local:    config.Local,
				SendType: config.SendType,
				Config:   config.Config,
				Level:    config.Level,
			},
		})
	}
	return nil
}

func (s *SendService) AddConfig(ctx context.Context, request *managepb.AddSendRequest) (*managepb.BaseResponse, error) {
	alert := model.AlertConfigPool.Get().(*model.AlertConfig)
	defer model.AlertConfigPool.Put(alert)
	parseAlertCheckByRequest(request, alert)
	err := s.dao.SaveAlertConfig(alert)
	if err != nil {
		return &managepb.BaseResponse{
			Code: managepb.ResponseCode_ERROR,
			Msg:  err.Error(),
		}, err
	}
	return &managepb.BaseResponse{
		Code: managepb.ResponseCode_SUCCESS,
		Msg:  "SUCCESS",
	}, nil
}

func (s *SendService) UpdateConfig(ctx context.Context, entry *managepb.SendEntry) (*managepb.BaseResponse, error) {
	alert := model.AlertConfigPool.Get().(*model.AlertConfig)
	defer model.AlertConfigPool.Put(alert)
	parseAlertCheckByEntry(entry, alert)
	err := s.dao.UpdateAlertConfig(alert)
	if err != nil {
		return &managepb.BaseResponse{
			Code: managepb.ResponseCode_ERROR,
			Msg:  err.Error(),
		}, err
	}
	return &managepb.BaseResponse{
		Code: managepb.ResponseCode_SUCCESS,
		Msg:  "SUCCESS",
	}, nil
}

func (s *SendService) DeleteConfig(ctx context.Context, request *managepb.IDRequest) (*managepb.BaseResponse, error) {
	err := s.dao.DelAlertConfigByID(request.ID)
	if err != nil {
		return &managepb.BaseResponse{
			Code: managepb.ResponseCode_ERROR,
			Msg:  err.Error(),
		}, err
	}
	return &managepb.BaseResponse{
		Code: managepb.ResponseCode_SUCCESS,
		Msg:  "SUCCESS",
	}, nil
}

func (s *SendService) GetConfigByID(ctx context.Context, request *managepb.IDRequest) (*managepb.SendConfigResponse, error) {
	config := s.dao.GetAlertConfigByID(request.ID)
	return &managepb.SendConfigResponse{
		Code: managepb.ResponseCode_SUCCESS,
		Msg:  "SUCCESS",
		Result: &managepb.SendEntry{
			ID:       config.ID,
			IP:       config.IP,
			Local:    config.Local,
			SendType: config.SendType,
			Config:   config.Config,
			Level:    config.Level,
		},
	}, nil
}

func (s *SendService) GetConfigsByAgent(request *managepb.AgentRequest, server managepb.SendService_GetConfigsByAgentServer) error {
	configs := s.dao.GetAlertConfigByIPAndLocal(request.IP, request.Local)
	for _, config := range configs {
		_ = server.Send(&managepb.SendConfigResponse{
			Code: managepb.ResponseCode_SUCCESS,
			Msg:  "SUCCESS",
			Result: &managepb.SendEntry{
				ID:       config.ID,
				IP:       config.IP,
				Local:    config.Local,
				SendType: config.SendType,
				Config:   config.Config,
				Level:    config.Level,
			},
		})
	}
	return nil
}

func NewSendService(s *setting.Setting) *SendService {
	return &SendService{
		dao: dao.NewStorageDao(s),
	}
}

func parseAlertCheckByRequest(request *managepb.AddSendRequest, alert *model.AlertConfig) {
	alert.IP = request.IP
	alert.Local = request.Local
	alert.Level = request.Level
	alert.SendType = request.SendType
	alert.Config = request.Config
}

func parseAlertCheckByEntry(entry *managepb.SendEntry, alert *model.AlertConfig) {
	alert.ID = entry.ID
	alert.IP = entry.IP
	alert.Local = entry.Local
	alert.Level = entry.Level
	alert.SendType = entry.SendType
	alert.Config = entry.Config
}
