package service

import (
	"context"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/dao"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/manage/managepb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
)

type AgentService struct {
	dao *dao.StorageDao
	*managepb.UnimplementedAgentServiceServer
}

func (a *AgentService) GetAllAgentInfo(request *managepb.BaseRequest, server managepb.AgentService_GetAllAgentInfoServer) error {
	agents := a.dao.GetAllAgentInfo()
	for _, agent := range agents {
		_ = server.Send(&managepb.AgentResponse{Code: managepb.ResponseCode_SUCCESS, Msg: "SUCCESS", Result: &managepb.AgentInfo{
			IP:      agent.IP,
			Local:   agent.Local,
			Port:    agent.Port,
			IsLive:  agent.IsLive,
			Metrics: agent.Metrics,
		}})
	}
	return nil
}

func (a *AgentService) GetAgentInfoByAgentID(ctx context.Context, request *managepb.AgentRequest) (*managepb.AgentResponse, error) {
	agent := a.dao.GetAgentInfoByIPAndLocal(request.IP, request.Local)
	return &managepb.AgentResponse{
		Code: managepb.ResponseCode_SUCCESS,
		Msg:  "SUCCESS",
		Result: &managepb.AgentInfo{
			IP:      agent.IP,
			Local:   agent.Local,
			Port:    agent.Port,
			IsLive:  agent.IsLive,
			Metrics: agent.Metrics,
		},
	}, nil
}

func NewAgentService(s *setting.Setting) *AgentService {
	return &AgentService{
		dao: dao.NewStorageDao(s),
	}
}
