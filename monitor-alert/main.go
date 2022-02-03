package main

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/global"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/pkg/setting"
)

func init() {
	setupSetting()
	setupService()
}

func main() {

}

func setupSetting() {
	global.Setting, _ = setting.NewSetting()
}

func setupService() {
	// 加载判定服务
	global.JudgementService = judgment.NewJudgmentService(global.Setting)
}
