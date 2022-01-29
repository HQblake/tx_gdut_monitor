package model

// AlertInfo 记录某台Agent的所有异常指标信息
type AlertInfo struct {
	AgentID string                // agent 标识
	Metrics map[string]MetricInfo // 指标告警信息
}

// MetricInfo 记录异常指标信息
type MetricInfo struct {
	Metric    string  // 指标类型
	Value     float64 // 指标值
	Threshold float64 // 阈值
	Method    int8    // 聚合方式
	Level     int8    // 告警等级
	Duration  string  // 持续时间
	Start     string  // 开始时间
}
