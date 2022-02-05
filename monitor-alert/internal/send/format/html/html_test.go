package html

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"testing"
	"time"
)

func TestHtml(t *testing.T) {
	info := model.Info{
		Agent: "127.0.0.1 广州",
		Metric: "CPU",
		Value:10,
		Threshold:3.14,
		Method:"Sum",
		Level:"panic",
		Duration:"5min",
		Start:time.Now().Format("[2006-01-01 15:04:05]"),
	}
	h := NewHtml()
	b, _ := h.Format(info)
	t.Log(string(b))
}
