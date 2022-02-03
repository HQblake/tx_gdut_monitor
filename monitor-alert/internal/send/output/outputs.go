package output

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"log"
	"sync"
)

type Config struct {
	Name   string
	Level  Level
	Config interface{}
}


type IOutputs interface {
	ID() string
	Get(id int) (IOutput, bool)
	Del(id int) error
	Set(id int, conf Config) error
	List() []IOutput
	Output(info model.MetricInfo) error
}

type Outputs struct {
	// agent id
	id string
	lock *sync.RWMutex
	// 对应agent的output集合
	outputs map[int]IOutput
}

func (o *Outputs) Output(info model.MetricInfo) error {
	var err error
	o.lock.RLock()
	for _, output := range o.outputs {
		if output.Level() <= Level(info.Level) {
			err = output.Output(info)
			log.Println(err)
		}
	}
	o.lock.RUnlock()
	return nil
}

func (o *Outputs) ID() string {
	return o.id
}

func (o *Outputs) List() []IOutput {
	o.lock.RLock()
	res := make([]IOutput, len(o.outputs))
	i := 0
	for _, k := range o.outputs {
		res[i] = k
		i++
	}
	o.lock.RUnlock()
	return res
}


func (o *Outputs) Get(id int) (IOutput, bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()
	output, ok := o.outputs[id]
	return output, ok
}

func (o *Outputs) Del(id int) error {
	o.lock.Lock()
	v, ok := o.outputs[id]
	if ok {
		err := v.Finish()
		if err != nil {
			return err
		}
		delete(o.outputs, id)
	}
	o.lock.Unlock()
	return nil
}

func (o *Outputs) Set(id int, conf Config) error {
	o.lock.Lock()
	defer o.lock.Unlock()
	v, ok := o.outputs[id]
	// 存在则更新配置
	if ok {
		err := v.Reset(conf.Level, conf.Config)
		if err != nil {
			return err
		}
		return nil
	}
	// 不存在则直接新增
	factory, err := Get(conf.Name)
	if err != nil {
		return err
	}
	c, err := ToConfig(conf.Config, factory.ConfigType())
	if err != nil {
		return err
	}
	output, err := factory.Create(conf.Level, c)
	if err != nil {
		return err
	}
	o.outputs[id] = output
	return nil
}

func NewOutputs(id string) *Outputs {
	return &Outputs{
		id: id,
		lock: &sync.RWMutex{},
		outputs: make(map[int]IOutput),
	}
}




type IManager interface {
	GetOutputs(id string) IOutputs
}
type Manager struct {
	lock *sync.RWMutex
	Agents map[string]IOutputs
}
func NewManager() *Manager {
	return &Manager{
		lock: &sync.RWMutex{},
		Agents: make(map[string]IOutputs),
	}
}

func (m *Manager) GetOutputs(id string) IOutputs {
	m.lock.RLock()
	o, ok := m.Agents[id]
	m.lock.RUnlock()
	if !ok {
		// 不存在则新增
		o = NewOutputs(id)
		m.lock.Lock()
		m.Agents[id] = o
		m.lock.Unlock()
	}
	return o
}