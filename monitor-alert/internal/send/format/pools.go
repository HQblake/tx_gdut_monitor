package format

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"sync"
)


var formatPool = newPools()

func Register(name string, pool IFormatterPool)  {
	formatPool.Register(name, pool)
}

func Get(name string) (IFormatterPool, error)  {
	return formatPool.Get(name)
}

func Format(name string, info model.Info) ([]byte, error) {
	pool, err := formatPool.Get(name)
	if err != nil {
		return nil, err
	}
	format := pool.Apply()
	defer pool.Release(format)
	return format.Format(info)
}


type IPoolManager interface {
	Register(name string, pool IFormatterPool)
	Get(name string) (IFormatterPool, error)
}

type Pools struct {
	lock *sync.RWMutex
	data map[string]IFormatterPool
}

func newPools() *Pools {
	return &Pools{
		data: make(map[string]IFormatterPool),
		lock: &sync.RWMutex{},
	}
}

func (p *Pools) Register(name string, pool IFormatterPool) {
	p.lock.Lock()
	p.data[name] = pool
	p.lock.Unlock()
}

func (p *Pools) Get(name string) (IFormatterPool, error) {
	p.lock.RLock()
	format, ok := p.data[name]
	p.lock.RUnlock()
	if !ok {
		return nil, fmt.Errorf("formatter %s is not register", name)
	}
	return format, nil
}


