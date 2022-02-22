package json

import (
	"encoding/json"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
)

type Json struct {

}

// Format json格式化直接输出
func (h *Json) Format(info []model.Info) ([]byte, error) {
	return json.Marshal(info)
}

func NewJson() *Json {
	return &Json{}
}

