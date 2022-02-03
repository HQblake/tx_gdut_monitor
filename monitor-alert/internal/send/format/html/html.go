package html

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
)

type Html struct {

}

func (h *Html) Format(alert model.MetricInfo) ([]byte, error) {
	panic("implement me")
}

func NewHtml() *Html {
	return &Html{}
}
