package line

import "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"

type Line struct {

}

func (h *Line) Format(alert model.MetricInfo) ([]byte, error) {
	panic("implement me")
}

func NewLine() *Line {
	return &Line{}
}

