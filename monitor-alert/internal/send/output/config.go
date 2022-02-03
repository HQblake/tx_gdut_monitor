package output

import (
	"encoding/json"
	"reflect"
)

func ToConfig(v interface{}, t reflect.Type) (interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	obj := newConfig(t)
	err = json.Unmarshal(data, obj)
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



