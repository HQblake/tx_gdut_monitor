package output

import "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"



// IOutput 输出器
type IOutput interface {
	Level() Level
	// Reset 更新配置
	Reset(level Level, config interface{}) error
	// Output 内容输出
	Output(info model.MetricInfo) error
	// Finish 结束
	Finish() error
}
