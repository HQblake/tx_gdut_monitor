package model

import "sync"

// AlertInfo 记录某台Agent的所有异常指标信息
type AlertInfo struct {
	IP      string                `json:"ip"`
	Local   string                `json:"local"`   // agent 标识
	Metrics map[string]MetricInfo `json:"metrics"` // 指标告警信息
}

// MetricInfo 记录异常指标信息
type MetricInfo struct {
	Metric    string  `json:"metric"`    // 指标类型
	Value     float64 `json:"value"`     // 指标值
	Threshold float64 `json:"threshold"` // 阈值
	Method    int32   `json:"method"`    // 聚合方式
	Level     int32   `json:"level"`     // 告警等级
	Duration  string  `json:"duration"`  // 持续时间
	Start     int64   `json:"start"`     // 开始时间
}

var AlertInfoPool = &sync.Pool{
	New: func() interface{} {
		alert := new(AlertInfo)
		alert.Metrics = make(map[string]MetricInfo)
		return alert
	},
}
