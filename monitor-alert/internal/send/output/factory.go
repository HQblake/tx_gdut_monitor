package output

import (
	"fmt"
	"reflect"
)


var factoryManager = newFactoryHandler()


type IFactoryManager interface {
	Register(name string, factory IFactory) error
	Get(name string) (IFactory, error)
}

func newFactoryHandler() *FactoryHandler {
	return &FactoryHandler{
		data: make(map[string]IFactory),
	}
}
func Get(name string) (IFactory, error) {
	return factoryManager.Get(name)
}

func Register(name string, factory IFactory) error {
	return factoryManager.Register(name, factory)
}

type FactoryHandler struct {
	data map[string]IFactory
}



func (f *FactoryHandler) Register(name string, factory IFactory) error {
	_, ok := f.data[name]
	if ok {
		return fmt.Errorf("duplicate register %s\n", name)
	}
	f.data[name] = factory
	return nil
}

func (f *FactoryHandler) Get(name string) (IFactory, error) {
	fac, ok := f.data[name]
	if !ok {
		return nil, fmt.Errorf("%s has not register \n", name)
	}
	return fac, nil
}

// IFactory 工厂
type IFactory interface {
	Create(level Level, config interface{}) (IOutput, error)
	ConfigType() reflect.Type
}

