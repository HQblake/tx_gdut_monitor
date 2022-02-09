package influxdb

import (
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Client struct {
	writeAPI api.WriteAPI
	queryAPI api.QueryAPI
	bucket   string
}

func NewClient(s *InfluxDBSetting) *Client {
	client := influxdb2.NewClientWithOptions(s.URL, s.Token,
		influxdb2.DefaultOptions().SetBatchSize(uint(s.BatchSize)))
	return &Client{client.WriteAPI(s.ORG, s.Bucket), client.QueryAPI(s.ORG), s.Bucket}
}
