package service

import (
	"context"
	"fmt"
	sendpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/api/gen"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strings"
)

type Service struct {
	agents output.IManager
	*grpc.Server
	*sendpb.UnimplementedSendServiceServer
}

func NewService(agents output.IManager, addr string) (*Service, error) {
	l, err := net.Listen("tcp",addr)
	if err != nil {
		return nil, err
	}
	s := &Service{
		agents: agents,
		Server: grpc.NewServer(),
	}
	sendpb.RegisterSendServiceServer(s.Server, s)
	go func() {
		err = s.Server.Serve(l)
		if err != nil {
			log.Println("grpc server:", err)
			return
		}
	}()
	return s, nil
}

func (s *Service) Set(ctx context.Context, request *sendpb.UpdateRequest) (*sendpb.Response, error) {
	outputs := s.agents.GetOutputs(fmt.Sprintf("%s-%s", request.GetIP(), request.GetLocal()))
	conf := output.Config{
		Name: strings.ToLower(request.GetConfig().GetConf().GetSendType().String()),
		Level: output.Level(request.GetConfig().GetConf().GetLevel()),
		Config: request.GetConfig().GetConf().GetConfig(),
	}
	err := outputs.Set(int(request.GetConfig().GetConfigID()), conf)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &sendpb.Response{
		Code: sendpb.ResponseCode_SUCCESS,
		Msg: "success",
	}, nil
}

func (s *Service) Del(ctx context.Context, request *sendpb.DelRequest) (*sendpb.Response, error) {
	outputs := s.agents.GetOutputs(fmt.Sprintf("%s-%s", request.GetIP(), request.GetLocal()))

	err := outputs.Del(int(request.GetConfigID()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &sendpb.Response{
		Code: sendpb.ResponseCode_SUCCESS,
		Msg: "success",
	}, nil
}

func (s *Service) Init(ctx context.Context, request *sendpb.InitRequest) (*sendpb.Response, error) {
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
				return nil, status.Errorf(codes.Internal, err.Error())
			}
		}
	}
	return &sendpb.Response{
		Code: sendpb.ResponseCode_SUCCESS,
		Msg: "success",
	}, nil
}



