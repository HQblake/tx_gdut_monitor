package agent

import (
	"context"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"
	managepb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/store/gen"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/send"
	"io"
	"log"
)

type Service struct {
	send send.ISend
	store managepb.AgentServiceClient
}

func (s *Service) GetAllAgentInfo() ([]model.AgentInfo, error) {
	var err error
	// 获取存储服务中对应agent的所有判定规则
	stream, err := s.store.GetAllAgentInfo(context.Background(), &managepb.BaseRequest{})
	if err != nil {
		return nil,  err
	}
	var resp *managepb.AgentResponse
	res := make([]model.AgentInfo, 0, 10)
	// 遍历获取指定metric的规则，没有则用默认规则代替
	for {
		resp, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("grpc get all agent info error %v", err)
			continue
		}
		if resp.Code != managepb.ResponseCode_SUCCESS{
			log.Printf("grpc get all agent info error %v", resp.Msg)
			continue
		}
		config := resp.GetResult()
		res = append(res,  model.AgentInfo{
			IP: config.GetIP(),
			Local: config.GetLocal(),
			Port: config.GetPort(),
			Metric: config.GetMetrics(),
			IsLive: config.IsLive,
		})
	}
	return res, nil
}

func (s *Service) GetAllAgentSendInfo() ([]model.AgentSendInfo, error) {
	var err error
	// 获取存储服务中对应agent的所有判定规则
	stream, err := s.store.GetAllAgentInfo(context.Background(), &managepb.BaseRequest{})
	if err != nil {
		return nil,  err
	}
	var resp *managepb.AgentResponse
	res := make([]model.AgentSendInfo, 0, 10)
	// 遍历获取指定metric的规则，没有则用默认规则代替
	for {
		resp, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("grpc get all agent info error %v", err)
			continue
		}
		if resp.Code != managepb.ResponseCode_SUCCESS{
			log.Printf("grpc get all agent info error %v", resp.Msg)
			continue
		}
		config := resp.GetResult()
		sendConfigs, err := s.send.GetConfigs(config.GetIP(), config.GetLocal())
		if err != nil {
			log.Printf("grpc get %s-%s send config error %v", config.GetIP(), config.GetLocal(), err)
			continue
		}
		sendType := make([]int32, 0, len(sendConfigs))
		for _, sendConfig := range sendConfigs {
			sendType = append(sendType, sendConfig.SendType)
		}
		res = append(res,  model.AgentSendInfo{
			IP: config.GetIP(),
			Local: config.GetLocal(),
			Port: config.GetPort(),
			IsLive: config.IsLive,
			SendConfig: sendType,
		})
	}
	return res, nil
}

func (s *Service) GetAgentInfo(ip string, local string) (*model.AgentInfo, error) {
	resp, err := s.store.GetAgentInfoByAgentID(context.Background(), &managepb.AgentRequest{IP: ip, Local: local})
	if err != nil {
		return nil, fmt.Errorf("grpc agent[%s-%s] info error %v",ip,local, err)
	}
	if resp.Code != managepb.ResponseCode_SUCCESS{
		return nil, fmt.Errorf("grpc agent[%s-%s] info error %v",ip,local, resp.Msg)
	}
	config := resp.GetResult()
	return &model.AgentInfo{
		IP: config.GetIP(),
		Local: config.GetLocal(),
		Port: config.GetPort(),
		Metric: config.GetMetrics(),
		IsLive: config.IsLive,
	},nil
}

func NewService(store managepb.AgentServiceClient, send send.ISend) *Service {
	return &Service{
		send: send,
		store: store,
	}
}