/*
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-13 22:53:14
 * @LastEditors: yzq
 */
package show

import (
	"time"

	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"
)

// IShow 图表展示的方案设计，可由自青设计
type IShow interface {
	// GetWarnInfo 返回当前时间节点之前所有告警信息
	GetWarnInfo() ([]model.HistoryInfo, error)
	// GetWarnInfoWithId 根据id返回告警信息
	GetWarnInfoWithId(id int32) ([]model.HistoryInfo, error)
	// GetWarnInfoWithParams 带参数查询告警信息, 参数为Ip和Local、Level、start、end
	GetWarnInfoWithParams(hinfo model.HistoryInfo, start, end time.Time) ([]model.HistoryInfo, error)

	// GetMetricsInOneDay 比如根据agentId(ip和local)和metric获取一天内的指标情况等
	GetMetricsWithTime(req model.MetricsReq, begin, end time.Time) ([]model.MetricsInfo, error)

	// DelWarnInfo 根据id删除告警信息
	DelWarnInfo(id int32) error
}
