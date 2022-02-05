package syslog

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"reflect"
)

func Register() {
	output.Register("syslog", NewFactory())
}

type factory struct {
	configType reflect.Type
}

func NewFactory() *factory {
	return &factory{
		configType: reflect.TypeOf(new(Config)),
	}
}

func (f *factory) Create(level output.Level, config interface{}) (output.IOutput, error) {
	panic("implement me")
}

func (f *factory) ConfigType() reflect.Type {
	return f.configType
}
