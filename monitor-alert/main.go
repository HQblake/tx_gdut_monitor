package main

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/global"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/pkg/setting"
	"log"
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
	var err error
	global.JudgementService, err = judgment.NewJudgmentService(global.Setting)
	if err != nil {
		log.Fatalln(err)
	}

	// 加载发送服务
	*global.SendService, err = send.NewService("")
	if err != nil {
		log.Fatalln(err)
	}
}
