package format

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
)

// IFormat 内容格式化工具，可以是json格式化，也可以自定义格式化方式
type IFormat interface {
	Format(info model.Info) ([]byte, error)
}

type IFormatterPool interface {
	Apply() (IFormat, bool)
	Release(format IFormat)
}




