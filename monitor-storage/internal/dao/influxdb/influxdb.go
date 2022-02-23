package influxdb

import (
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"log"
	"time"
)

type Client struct {
	writeAPI api.WriteAPI
	queryAPI api.QueryAPI
	bucket   string
}

func NewClient(s *InfluxDBSetting) *Client {
	// 定义选项
	option := influxdb2.DefaultOptions()
	option.SetLogLevel(uint(s.LogLevel))
	option.SetBatchSize(uint(s.BatchSize))
	option.SetFlushInterval(uint(s.FlushIntervalMs))
	option.SetRetryInterval(uint(s.RetryInterval))
	option.SetMaxRetryInterval(uint(s.MaxRetryIntervalMs))
	option.SetMaxRetries(uint(s.MaxRetries))
	option.SetMaxRetryTime(uint(s.MaxRetryTime))
	option.SetPrecision(time.Duration(s.Precision) * time.Nanosecond)
	option.SetExponentialBase(uint(s.ExponentialBase))
	option.SetHTTPRequestTimeout(uint(s.HttpRequestTimeout))
	option.SetRetryBufferLimit(uint(s.RetryBufferLimit))
	option.SetUseGZip(s.UseGZip)
	client := influxdb2.NewClientWithOptions(s.URL, s.Token,
		option)
	defer log.Println("InfluxDB Connection Succeeded")
	return &Client{client.WriteAPI(s.ORG, s.Bucket), client.QueryAPI(s.ORG), s.Bucket}
}
