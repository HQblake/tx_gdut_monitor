package influxdb

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	"log"
	"math/rand"
	"testing"
	"time"
)

var setting = &InfluxDBSetting{
	URL:                "http://127.0.0.1:8086",
	ORG:                "personal",
	Token:              "FTGYhVo9DwF5v56ZxFC4j56R4TE-L6scyeO5jBf-eHsKPjj7O1dRl5dz2sDf65DuKUl-vlvt3exVrzuMBWKzig==",
	Bucket:             "metric_test",
	BatchSize:          20,
	LogLevel:           0,
	FlushIntervalMs:    1000,
	RetryInterval:      5000,
	MaxRetryIntervalMs: 125000,
	MaxRetries:         5,
	MaxRetryTime:       180000,
	ExponentialBase:    2,
	UseGZip:            false,
	RetryBufferLimit:   50000,
	Precision:          1,
	HttpRequestTimeout: 20,
}

func TestClient_SaveMatricData(t *testing.T) {
	client := NewClient(setting)
	writeNum := 10000

	// 创建时间戳
	now := time.Now().Unix()
	times := make([]int64, 0, writeNum)
	for i := 0; i < writeNum; i++ {
		times = append(times, now+int64(i))
	}

	for i := 0; i < writeNum; i++ {
		client.SaveMatricData(&model.Metric{
			Name:       "cpu_rate",
			Value:      rand.Float64(),
			IP:         "127.0.0.1",
			Local:      "上海",
			Port:       "8080",
			Timestamp:  times[i],
			Dimensions: make(map[string]string),
		})
	}
	fmt.Println(times[0], " ~ ", times[writeNum-1])
}

func TestClient_GetMetricData(t *testing.T) {
	client := NewClient(setting)
	metrics, err := client.GetMetricData("127.0.0.1", "上海", "cpu_rate", "100s",
		1644491401, 1644501401, 2, 30)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(len(metrics))
}

func TestClient_GetAggregatedData(t *testing.T) {
	client := NewClient(setting)
	result, err := client.GetAggregatedData(&model.Metric{
		Name:  "cpu_rate",
		IP:    "127.0.0.1",
		Local: "上海",
	}, "10s", 8, 1644501401)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}
