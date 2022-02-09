package model

import "sync"

type Metric struct {
	Name       string            `json:"name"`
	Value      float64           `json:"value"`
	IP         string            `json:"ip"`
	Local      string            `json:"local"`
	Port       string            `json:"port"`
	Timestamp  int64             `json:"timestamp"`
	Dimensions map[string]string `json:"dimensions"`
}

var MetricPool = sync.Pool{
	New: func() interface{} {
		return new(Metric)
	},
}
