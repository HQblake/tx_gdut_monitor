package service

import (
	"context"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/dao"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/manage/managepb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
)

type JudgmentService struct {
	dao *dao.StorageDao
	*managepb.UnimplementedJudgmentServiceServer
}

func (j *JudgmentService) GetConfigsByAgent(request *managepb.AgentRequest, server managepb.JudgmentService_GetConfigsByAgentServer) error {
	configs := j.dao.GetCheckConfigsByIPAndLocal(request.IP, request.Local)
	for _, config := range configs {
		_ = server.Send(&managepb.JudgmentConfigResponse{
			Code: managepb.ResponseCode_SUCCESS,
			Msg:  "SUCCESS",
			Result: &managepb.JudgmentEntry{
				ID:        config.ID,
				IP:        config.IP,
				Local:     config.Local,
				Metric:    config.Metric,
				Method:    config.Method,
				Period:    config.Period,
				Threshold: config.Threshold,
			},
		})
	}
	return nil
}

func (j *JudgmentService) UpdateConfig(ctx context.Context, entry *managepb.JudgmentEntry) (*managepb.BaseResponse, error) {
	check := model.CheckConfigPool.Get().(*model.CheckConfig)
	defer model.CheckConfigPool.Put(check)
	parseCheckConfig(entry, check)
	id, err := j.dao.UpdateCheckConfig(check)
	if err != nil {
		return &managepb.BaseResponse{
			Code: managepb.ResponseCode_ERROR,
			Msg:  err.Error(),
		}, err
	}
	return &managepb.BaseResponse{
		Code: managepb.ResponseCode_SUCCESS,
		Msg:  string(id),
	}, nil
}

func (j *JudgmentService) DeleteConfig(ctx context.Context, request *managepb.IDRequest) (*managepb.BaseResponse, error) {
	err := j.dao.DelCheckConfigByID(request.ID)
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

func NewJudgmentService(s *setting.Setting) *JudgmentService {
	return &JudgmentService{
		dao: dao.NewStorageDao(s),
	}
}

func parseCheckConfig(entry *managepb.JudgmentEntry, check *model.CheckConfig) {
	check.ID = entry.ID
	check.IP = entry.IP
	check.Local = entry.Local
	check.Metric = entry.Metric
	check.Period = entry.Period
	check.Method = entry.Method
	check.Threshold = entry.Threshold
}
