package setting

import (
	"fmt"
	"log"
	"testing"
)

type mySQLSetting struct {
	DBType          string
	Username        string
	Password        string
	Host            string
	DBName          string
	Charset         string
	ParseTime       bool
	MaxIdleConns    int
	MaxOpenConns    int
	MaxLifetime     int
	SingularTable   bool
	SkipTransaction bool
}

type influxDBSetting struct {
	URL                string
	ORG                string
	Token              string
	Bucket             string
	LogLevel           int
	BatchSize          int
	FlushIntervalMs    int
	RetryInterval      int
	MaxRetryIntervalMs int
	MaxRetries         int
	MaxRetryTime       int
	ExponentialBase    int
	UseGZip            bool
	RetryBufferLimit   int
	Precision          int
	HttpRequestTimeout int
}

// 需要先将 monitor-storage/configs 目录拷贝到当前目录下
func TestSetting_ReadSection(t *testing.T) {
	setting, err := NewSetting()
	if err != nil {
		log.Fatalln(err)
	}
	ms := &mySQLSetting{}
	is := &influxDBSetting{}
	err = setting.ReadSection("MySQL", ms)
	if err != nil {
		log.Fatalln(err)
	}
	err = setting.ReadSection("InfluxDB", is)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ms)
	fmt.Println(is)
}
