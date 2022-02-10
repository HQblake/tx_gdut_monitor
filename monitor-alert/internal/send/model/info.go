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
var methodType = map[int32]string{
	0: "计数",
	1: "总和",
	2: "平均值",
	3: "中位数",
	4: "积分",
	5: "众数",
	6: "极值",
	7: "标准差",
	8: "最大值",
	9: "最小值",
}
// ParseMethod 解析聚合方式
func (i *Info) ParseMethod(method int32) {
	m, ok := methodType[method]
	if ok {
		i.Method = m
	}
	i.Method = "unknown"
}
