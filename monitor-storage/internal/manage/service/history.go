/*
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-11 10:35:39
 * @LastEditors: yzq
 */
package service

import (
	"context"

	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/dao"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/manage/managepb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
)

type HistoryService struct {
	dao *dao.StorageDao
	*managepb.UnimplementedHistoryServiceServer
}

func (h *HistoryService) GetAllAlertInfo(request *managepb.BaseRequest, server managepb.HistoryService_GetAllAlertInfoServer) error {
	histories := h.dao.GetAllAlertInfo()
	for _, info := range histories {
		_ = server.Send(&managepb.AlertResponse{
			Code: managepb.ResponseCode_SUCCESS,
			Msg:  "SUCCESS",
			Result: &managepb.AlertInfo{
				ID:        info.ID,
				IP:        info.IP,
				Local:     info.Local,
				Metric:    info.Metric,
				Value:     info.Value,
				Threshold: info.Threshold,
				Duration:  info.Duration,
				Level:     info.Level,
				Start:     info.Start,
				Method:    info.Method,
			},
		})
	}
	return nil
}

func (h *HistoryService) GetAlertInfo(request *managepb.AlertRequest, server managepb.HistoryService_GetAlertInfoServer) error {
	histories := h.dao.GetAlertInfo(request.ID, request.Level, request.IP, request.Local, request.Metric, request.Begin, request.End)
	for _, info := range histories {
		_ = server.Send(&managepb.AlertResponse{
			Code: managepb.ResponseCode_SUCCESS,
			Msg:  "SUCCESS",
			Result: &managepb.AlertInfo{
				ID:        info.ID,
				IP:        info.IP,
				Local:     info.Local,
				Metric:    info.Metric,
				Value:     info.Value,
				Threshold: info.Threshold,
				Duration:  info.Duration,
				Level:     info.Level,
				Start:     info.Start,
				Method:    info.Method,
			},
		})
	}
	return nil
}

func (h *HistoryService) DelAlterInfo(ctx context.Context, request *managepb.IDRequest) (*managepb.BaseResponse, error) {
	err := h.dao.DelAlterInfo(request.ID)
	if err != nil {
		return &managepb.BaseResponse{Code: managepb.ResponseCode_ERROR, Msg: err.Error()}, err
	}
	return &managepb.BaseResponse{Code: managepb.ResponseCode_SUCCESS, Msg: "SUCCESS"}, nil
}

func NewHistoryService(s *setting.Setting) *HistoryService {
	return &HistoryService{
		dao: dao.NewStorageDao(s),
	}
}
