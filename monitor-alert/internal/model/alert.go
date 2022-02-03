package model

import "sync"

// AlertInfo 记录某台Agent的所有异常指标信息
type AlertInfo struct {
	AgentID string                `json:"agent_id"` // agent 标识
	Metrics map[string]MetricInfo `json:"metrics"`  // 指标告警信息
}

// MetricInfo 记录异常指标信息
type MetricInfo struct {
	Metric    string  `json:"metric"`    // 指标类型
	Value     float64 `json:"value"`     // 指标值
	Threshold float64 `json:"threshold"` // 阈值
	Method    int8    `json:"method"`    // 聚合方式
	Level     int8    `json:"level"`     // 告警等级
	Duration  string  `json:"duration"`  // 持续时间
	Start     string  `json:"start"`     // 开始时间
}

var AlertInfoPool = &sync.Pool{
	New: func() interface{} {
		return new(AlertInfo)
	},
}
