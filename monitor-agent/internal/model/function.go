package model

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-agent/pkg/dimensions"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-agent/pkg/metrics"
)

var GlobalMetrics = map[string]func() float64{ //全局的指标函数Map
	"cpu_rate": metrics.GetCpuRate,
	"mem_rate": metrics.GetMemRate,
}
var globalDimensions = map[string]func() string{ //全局的维度函数Map
	"ip_address": dimensions.GetIP,
	"port":       dimensions.GetPort,
}
