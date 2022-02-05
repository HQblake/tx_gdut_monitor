package output

import (
	"encoding/json"
	"reflect"
	"testing"
)

type Test struct {
	Name string `json:"name"`
}

func TestToConfig(t *testing.T) {
	test := &Test{
		Name: "test",
	}
	data, err := json.Marshal(test)
	if err != nil {
		t.Errorf(err.Error())
	}
	conf, err := ToConfig(string(data), reflect.TypeOf(new(Test)))
	if err != nil {

	}
	if !reflect.DeepEqual(test, conf) {
		t.Errorf("parse config error, want %+v, but got %+v", test, conf)
	}
}