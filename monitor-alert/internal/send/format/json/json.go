package json

import "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"

type Json struct {

}

func (h *Json) Format(alert model.MetricInfo) ([]byte, error) {
	panic("implement me")
}

func NewJson() *Json {
	return &Json{}
}

