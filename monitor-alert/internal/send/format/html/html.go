package html

import "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"

type Html struct {

}

func (h *Html) Format(info model.Info) ([]byte, error) {

	panic("implement me")
}

func NewHtml() *Html {
	return &Html{}
}
