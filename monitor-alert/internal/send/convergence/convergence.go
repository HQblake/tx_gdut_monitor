package convergence

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/global/setting"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/convergence/aggregation"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/convergence/instant"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/convergence/roll"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"log"
	"time"
)

type IConvergence interface {
	Alert(alert *model.AlertInfo) error
}

func NewConvergence(manage output.IManager) IConvergence {
	config, err := setting.GetAlertConfig()
	if err != nil {
		log.Fatal(err)
	}
	if config.Interval == 0 {
		config.Interval = 1
	}
	switch config.Convergence {
	case 1:
		// 简单聚合收敛处理，默认一分钟告警一次
		return aggregation.NewConvergence(manage, time.Duration(config.Interval) * time.Second)
	case 2:
		// 滚动收敛聚合处理，默认一分钟告警一次
		return roll.NewConvergence(manage, time.Duration(config.Interval) * time.Second)
	default:
		// 默认为0，不做收敛处理，收到即刻告警
		return instant.NewConvergence(manage)
	}
}