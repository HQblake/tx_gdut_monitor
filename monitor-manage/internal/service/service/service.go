package service

import (
	"context"
	"encoding/json"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/configs"
	judgpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/service/gen"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/judgment"
	"log"
)

type Service struct {
	judgment judgment.IJudgment
	*judgpb.UnimplementedManageServiceServer
}

func NewService(judgment judgment.IJudgment) *Service {
	return &Service{
		judgment: judgment,
	}
}

// Get 配合check服务的GET需求
func (s *Service) Get(ctx context.Context, request *judgpb.CheckRequest) (*judgpb.CheckResponse, error) {
	// 获取存储服务中对应agent的所有判定规则
	res := make(map[string]*judgpb.MetricRule, len(request.GetMetrics()))
	_, cfgs, err := s.judgment.GetConfigs(request.GetIP(), request.GetLocal())
	if err != nil {
		return &judgpb.CheckResponse{
			Code: judgpb.CheckResponse_ERROR,
			Msg: err.Error(),
		},nil
	}
	defaultRule := configs.GetDefaultRule()
	for _, m := range request.GetMetrics() {
		if _, ok := cfgs[m]; ok {
			var threshold map[int32]float64
			err = json.Unmarshal([]byte(cfgs[m].Threshold), &threshold)
			if err != nil {
				log.Printf("grpc get judgment json parse threshold rule error %v", err)
				continue
			}
			res[m] = &judgpb.MetricRule{
				Method: cfgs[m].Method,
				Period: cfgs[m].Period,
				Threshold: threshold,
			}
		}else {
			res[m] = &judgpb.MetricRule{
				Method: defaultRule.Method,
				Period: defaultRule.Period,
				Threshold: defaultRule.Threshold,
			}
		}
	}
	return &judgpb.CheckResponse{
		Code: judgpb.CheckResponse_SUCCESS,
		Msg: "success",
		Result: &judgpb.AgentRule{
			IP: request.GetIP(),
			Local: request.GetLocal(),
			Metrics: res,
		},
	},nil
}



