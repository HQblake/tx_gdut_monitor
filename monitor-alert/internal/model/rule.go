package model

// AgentRule 记录某台Agent所有指标的判定规则
type AgentRule struct {
	AgentID string                `json:"agent_id"`     // Agent 唯一标识
	Metrics map[string]MetricRule `json:"metrics_rule"` // 所有指标的判定规则
}

// MetricRule 记录指标的判定规则
type MetricRule struct {
	Method    int8            `json:"type"`      // 指标的聚合方式
	Period    string          `json:"period"`    // 指标的聚合周期 eg. 格式如：5s、5m、5h、1h1m1s....
	Threshold map[int]float64 `json:"threshold"` // 阈值等级判定
}
