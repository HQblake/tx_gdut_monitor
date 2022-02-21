package convergence

import "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"

type IConvergence interface {
	Alert(alert *model.AlertInfo) error
}