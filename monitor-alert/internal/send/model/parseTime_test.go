package model

import "testing"

func TestParseTime(t *testing.T) {
	i := Info{}
	i.ParseDate(1645686032)
	t.Log(i.Start)
}
