/*
 * @Description:
 * @Autor: yzq
 * @Date: 2022-02-23 13:52:42
 * @LastEditors: yzq
 */
package influxdb

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
)

var setting = &InfluxDBSetting{
	URL:                "http://81.71.165.211:8086/",
	ORG:                "tx-monitor",
	Token:              "W494CW95RhZzEc95DpxubUw19EoBuKxCxAlv3YfVCS4KRmYnJf6MhPJGWXFvzkX4hJ76-ZDBIwpV2D76wIZMcA==",
	Bucket:             "metric",
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
	writeNum := 1

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
	begin, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-02-21 00:00:00", time.UTC)
	start := begin.Unix()
	end, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-02-22 00:00:00", time.UTC)
	stop := end.Unix()
	metrics, err := client.GetMetricData("172.18.0.5", "beijing", "cpu_rate", "120s",
		start, stop, 0, 0)
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
