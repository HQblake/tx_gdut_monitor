package dao

import "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"

func GetAggregatedData(period string, method int32, metric *model.Metric) (float64, error) {
	return 0, nil
}

func SetAgentInfo(metric *model.Metric) error {
	return nil
}

func SaveAlertInfo(alert *model.AlertInfo) error {
	return nil
}

func GetAllAgentInfo() []model.AgentInfo {
	return nil
}

func GetMetricsByAgentID(agentID string) []string {
	return nil
}

func GetMetricData(agentID, metricName, period string, begin, end int64) []model.Metric {
	return nil
}

func GetAllAlertInfo() []model.AlertInfo {
	return nil
}

func GetAlertInfo(id, level int32, agentId, metric string, begin, end int64) []model.AlertInfo {
	return nil
}

func DelAlterInfo(id int32) error {
	return nil
}
