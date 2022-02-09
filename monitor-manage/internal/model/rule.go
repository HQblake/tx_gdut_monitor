package model

import "sync"

// AgentRule 记录某台Agent所有指标的判定规则
type AgentRule struct {
	AgentID string                `json:"agent_id"`     // Agent 唯一标识
	Metrics map[string]MetricRule `json:"metrics_rule"` // 所有指标的判定规则
}

// MetricRule 记录指标的判定规则
type MetricRule struct {
	Method    int8             `json:"method"`      // 指标的聚合方式
	Period    string           `json:"period"`    // 指标的聚合周期 eg. 格式如：5s、5m、5h、1h1m1s....
	Threshold map[int8]float64 `json:"threshold"` // 阈值等级判定
}

var AgentRulePool = &sync.Pool{
	New: func() interface{} {
		return new(AgentRule)
	},
}

type CheckConfig struct {
	Id int `json:"id"`
	Ip string `json:"ip"`
	Local string `json:"local"`
	Metric string `json:"metric"`
	Method int `json:"method"`
	Period string `json:"period"`
	Threshold string `json:"threshold"`
}