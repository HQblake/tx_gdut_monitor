package json

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format"
	"sync"
)

const name = "json"

func Register() {
	format.Register(name, NewPool())
}

type pool struct {
	p *sync.Pool
}

func (f *pool) Apply() (format.IFormat, bool) {
	formatter, ok := f.p.Get().(format.IFormat)
	return formatter, ok

}

func (f *pool) Release(format format.IFormat) {
	f.p.Put(format)
}

func NewPool() *pool {
	return &pool{
		p : &sync.Pool{
			New: func() interface{} {
				return NewJson()
			},
		},
	}
}
