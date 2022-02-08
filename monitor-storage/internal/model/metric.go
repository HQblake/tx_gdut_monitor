package model

import "sync"

type Metric struct {
	Name       string
	Value      float64
	IP         string
	Local      string
	Port       string
	Timestamp  int64
	Dimensions map[string]string
}

var MetricPool = sync.Pool{
	New: func() interface{} {
		return new(Metric)
	},
}
