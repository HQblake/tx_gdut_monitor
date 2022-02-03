package model

import "sync"

// AgentReport Agent采集信息上报的数据结构
type AgentReport struct {
	IP         string             // Agent的IP地址
	Port       string             // Agent的端口号
	Local      string             // Agent的区域
	Timestamp  int64              // 信息上报的时间戳
	Metrics    map[string]float64 // 指标类型: 指标值
	Dimensions map[string]string  // Agent维度
}

var AgentReportPool = &sync.Pool{
	New: func() interface{} {
		return new(AgentReport)
	},
}