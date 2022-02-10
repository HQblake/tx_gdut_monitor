package model

type MetricsInfo struct {
	Timestamp int64 `json:"timestamp" yaml:"timestamp"`
	Metric string `json:"metric" yaml:"metric"`
	Value float64 `json:"value" yaml:"value"`
}
