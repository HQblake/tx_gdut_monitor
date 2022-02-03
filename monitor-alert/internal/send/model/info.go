package model

// Info 格式化结构
type Info struct {
	Agent     string  `json:"agent"`
	Metric    string  `json:"metric"`    // 指标类型
	Value     float64 `json:"value"`     // 指标值
	Threshold float64 `json:"threshold"` // 阈值
	Method    string  `json:"method"`    // 聚合方式
	Level     string  `json:"level"`     // 告警等级
	Duration  string  `json:"duration"`  // 持续时间
	Start     string  `json:"start"`     // 开始时间
}

// ParseMethod 解析聚合方式
// todo int类型，需要再进行解析
func (i *Info) ParseMethod(method int8) {
	i.Method = string(method)
}
