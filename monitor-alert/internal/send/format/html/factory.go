package html

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/format"
	"sync"
)

const name = "html"

func Register() {
	format.Register(name, NewPool())
}

type pool struct {
	p *sync.Pool
}

func (f *pool) Apply() format.IFormat {
	return f.p.Get().(format.IFormat)
}

func (f *pool) Release(format format.IFormat) {
	f.p.Put(format)
}

func NewPool() *pool {
	return &pool{
		p : &sync.Pool{
			New: func() interface{} {
				return NewHtml()
			},
		},
	}
}



