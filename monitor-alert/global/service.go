package global

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/receive"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send"
)

// 对告警系统内部的所有服务模块进行注册
var (
	JudgementService *judgment.JudgmentService
	SendService      *send.SendService
	ReceiveService   *receive.ReceiveService
)
