package show

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"
)

// IShow 图表展示的方案设计，可由自青设计
type IShow interface {
	// GetMetricWithTime 根据agentId(ip和local)和时间起始以及时间长度获取一天内的指标情况
	// GetMetricWithTime(ip string, local string, timeStamp int64, counter int) ([]model.MetricsInfo, error)
	// GetWarnInfo 返回当前时间节点之前所有告警信息
	GetWarnInfo() ([]model.HistoryInfo, error)
	// GetWarnInfoWithLevel 按等级返回所有告警信息
	GetWarnInfoWithLevel(level int32) ([]model.HistoryInfo, error)
	// GetWarnInfoWithTimestamp 按时间返回所有告警信息
	GetWarnInfoWithTimestamp(timeStamp int64) ([]model.HistoryInfo, error)
	// GetWarnInfoWithId 根据agentId(ip和local)返回所有告警信息
	GetWarnInfoWithId(ip string, local string) ([]model.HistoryInfo, error)

	// GetMetricsInOneDay 比如根据agentId(ip和local)和metric获取一天内的指标情况等
	GetMetricsWithTime(ip string, local string, metric string, begin int64, limit int32) ([]model.MetricsInfo, error)
}
