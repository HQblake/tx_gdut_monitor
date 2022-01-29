package model

// AgentReport Agent采集信息上报的数据结构
type AgentReport struct {
	IP        string             // Agent的IP地址
	port      string             // Agent的端口号
	Local     string             // Agent的区域
	Timestamp int                // 信息上报的时间戳
	Metrics   map[string]float64 // 指标类型: 指标值
}
