package mail

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"reflect"
)

func Register() {
	output.Register("email", NewFactory())
}

type factory struct {

}

func NewFactory() *factory {
	return &factory{}
}

func (f *factory) Create(level output.Level, config interface{}) (output.IOutput, error) {
	panic("implement me")
}

func (f *factory) ConfigType() reflect.Type {
	panic("implement me")
}
