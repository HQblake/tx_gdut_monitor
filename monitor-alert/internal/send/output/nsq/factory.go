package nsq

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"reflect"
)

func Register() {
	output.Register("nsq", NewFactory())
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
	conf, ok := config.(*Config)
	if !ok {
		return nil, fmt.Errorf("config type is invalid")
	}
	err := conf.doCheck()
	if err != nil {
		return nil, err
	}
	return NewNsq(level, conf)
}

func (f *factory) ConfigType() reflect.Type {
	return f.configType
}
