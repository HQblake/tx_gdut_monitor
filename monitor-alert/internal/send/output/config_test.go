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
	test := Test{
		Name: "test",
	}
	data, err := json.Marshal(test)
	if err != nil {
		t.Errorf(err.Error())
	}
	conf, err := ToConfig(string(data), reflect.TypeOf(new(Test)))
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(conf)




	//type args struct {
	//	v string
	//	t reflect.Type
	//}
	//tests := []struct {
	//	name    string
	//	args    args
	//	want    interface{}
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		got, err := ToConfig(tt.args.v, tt.args.t)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("ToConfig() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("ToConfig() got = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}