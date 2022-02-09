package global

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/global/service"
	set "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
)

var (
	Setting         *set.Setting
	JudgmentService *service.JudgmentService
	ManageService   *service.ManageService
)
