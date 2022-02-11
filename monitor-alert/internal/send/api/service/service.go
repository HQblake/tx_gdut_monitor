package service

import (
	"context"
	"fmt"
	sendpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/api/gen"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"log"
	"strings"
)

type Service struct {
	agents output.IManager
	*sendpb.UnimplementedSendServiceServer
}

func NewService(agents output.IManager) *Service {
	return &Service{
		agents: agents,
	}
}

func (s *Service) Set(ctx context.Context, request *sendpb.UpdateRequest) (*sendpb.SendResponse, error) {
	outputs := s.agents.GetOutputs(fmt.Sprintf("%s-%s", request.GetIP(), request.GetLocal()))
	conf := output.Config{
		Name: strings.ToLower(request.GetConfig().GetConf().GetSendType().String()),
		Level: output.Level(request.GetConfig().GetConf().GetLevel()),
		Config: request.GetConfig().GetConf().GetConfig(),
	}
	err := outputs.Set(int(request.GetConfig().GetConfigID()), conf)
	if err != nil {
		return &sendpb.SendResponse{
			Code: sendpb.SendResponse_ERROR,
			Msg: err.Error(),
		}, nil
	}
	return &sendpb.SendResponse{
		Code: sendpb.SendResponse_SUCCESS,
		Msg: "success",
	}, nil
}

func (s *Service) Del(ctx context.Context, request *sendpb.DelRequest) (*sendpb.SendResponse, error) {
	outputs := s.agents.GetOutputs(fmt.Sprintf("%s-%s", request.GetIP(), request.GetLocal()))

	err := outputs.Del(int(request.GetConfigID()))
	if err != nil {
		return &sendpb.SendResponse{
			Code: sendpb.SendResponse_ERROR,
			Msg: err.Error(),
		}, nil
	}
	return &sendpb.SendResponse{
		Code: sendpb.SendResponse_SUCCESS,
		Msg: "success",
	}, nil
}

func (s *Service) Init(ctx context.Context, request *sendpb.InitRequest) (*sendpb.SendResponse, error) {
	for agent, config := range request.GetConfig() {
		outputs := s.agents.GetOutputs(agent)
		for _, c := range config.GetConfig() {
			conf := output.Config{
				Name: strings.ToLower(c.GetConf().GetSendType().String()),
				Level: output.Level(c.GetConf().GetLevel()),
				Config: c.GetConf().GetConfig(),
			}
			err := outputs.Set(int(c.GetConfigID()), conf)
			if err != nil {
				// 打日志，不影响其他
				log.Printf("init agent %s, set config id %d error:%s", agent, c.GetConfigID(), err.Error())
			}
		}
	}
	return &sendpb.SendResponse{
		Code: sendpb.SendResponse_SUCCESS,
		Msg: "success",
	}, nil
}



