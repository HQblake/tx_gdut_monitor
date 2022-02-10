package mysql

import (
	"log"
	"testing"
)

var setting = &MySQLSetting{
	DBType:       "mysql",
	Username:     "root",
	Password:     "root",
	Host:         "127.0.0.1:3306",
	DBName:       "monitor",
	Charset:      "utf8mb4",
	ParseTime:    true,
	MaxIdleConns: 10,
	MaxOpenConns: 10,
	MaxLifetime:  3,
}

func TestNewClient(t *testing.T) {
	client := NewClient(setting)
	_, err := client.db.Exec("CALL AddAgentInfo(?,?,?,?)",
		"127.0.0.1", "上海", "8080", "cpu_rate")
	if err != nil {
		log.Fatalln(err)
	}
}
