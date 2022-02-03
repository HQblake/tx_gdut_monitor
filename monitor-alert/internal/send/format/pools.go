package format

import (
	"fmt"
	"sync"
)


type IFormatterPool interface {
	Apply() (IFormat, bool)
	Release(format IFormat)
}

type IPoolManager interface {
	Register(name string, pool IFormatterPool)
	Get(name string) (IFormatterPool, error)
}


var formatPool = newPools()

func Register(name string, pool IFormatterPool)  {
	formatPool.Register(name, pool)
}

func Get(name string) (IFormatterPool, error)  {
	return formatPool.Get(name)
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


