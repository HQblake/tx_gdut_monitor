package service

import (
	"context"
	sendpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/api/gen"
)

type Service struct {
	*sendpb.UnimplementedSendServiceServer
}

func (s *Service) Update(ctx context.Context, request *sendpb.UpdateRequest) (*sendpb.Response, error) {
	panic("implement me")
}

func (s *Service) Init(ctx context.Context, request *sendpb.InitRequest) (*sendpb.Response, error) {
	panic("implement me")
}


