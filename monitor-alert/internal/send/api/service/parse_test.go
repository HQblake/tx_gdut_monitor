package service

import (
	sendpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/api/gen"
	"testing"
)

func TestSendType(t *testing.T) {
	t.Log(sendpb.Type(0).String())
}
