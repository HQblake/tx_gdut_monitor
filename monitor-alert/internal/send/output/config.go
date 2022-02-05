package output

import (
	"encoding/json"
	"reflect"
)

func ToConfig(v string, t reflect.Type) (interface{}, error) {
	obj := newConfig(t)
	err := json.Unmarshal([]byte(v), obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func newConfig(p reflect.Type) interface{} {
	if p.Kind() == reflect.Ptr {
		p = p.Elem()
	}
	return reflect.New(p).Interface()
}



