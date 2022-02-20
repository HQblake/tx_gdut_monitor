package model

type MetricsInfo struct {
	Timestamp int64   `json:"timestamp" yaml:"timestamp"`
	Metric    string  `json:"metric" yaml:"metric"`
	Value     float64 `json:"value" yaml:"value"`
}

type MetricsReq struct {
	IP         string `json:"ip"`
	Local      string `json:"local"`
	MetricName string `json:"metricName"`
	Period     string `json:"period"`
	Begin      string `json:"begin"`
	End        string `json:"end"`
	Method     string `json:"method"`
	Limit      int    `json:"limit"`
}
