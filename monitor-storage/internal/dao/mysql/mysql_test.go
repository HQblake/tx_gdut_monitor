package mysql
import "testing"

func TestAdd(t *testing.T) {
	s:=MySQLSetting{
		"",
		"root",
		"",
		"",
		"huangtest",
		"",
		true,
		10,
		5,
		0,
		true,
		true,
	}
	NewClient(&s)
}
