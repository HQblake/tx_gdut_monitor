package influxdb

import (
	"context"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

// SaveMatricData 将上报记录保存到InfluxDB中
func (c *Client) SaveMatricData(metric *model.Metric) error {
	tags, fields := parseTagsAndFileds(metric)
	c.writeAPI.WritePoint(influxdb2.NewPoint(metric.Name, tags, fields, time.Unix(metric.Timestamp, 0)))
	c.writeAPI.Flush()
	return nil
}

func (c *Client) GetAggregatedData(metric *model.Metric, period string, method int32, timestamp int64) (float64, error) {
	query := `from(bucket: "%s")
			  |> range(start: %v, stop: %v)
			  |> filter(fn: (r) => r["_measurement"] == "%s"
			  |> filter(fn: (r) => r["ip"] == "%s")
			  |> filter(fn: (r) => r["local"] == "%s")
			  |> filter(fn: (r) => r["_field"] == "value")
			  |> aggregateWindow(every: %s, fn: %s, createEmpty: false)`
	duration, err := time.ParseDuration(period)
	if err != nil {
		return 0, err
	}
	stop := time.Unix(timestamp, 0).Add(1 * time.Second)
	start := stop.Add(-(duration + 1))

	result, err := c.queryAPI.Query(context.Background(),
		fmt.Sprintf(query, c.bucket, start, stop, metric.Name, metric.IP, metric.Local, period, Methods[method]))
	if err != nil {
		return 0, err
	}

	if result.Next() {
		return result.Record().Value().(float64), nil
	} else {
		return metric.Value, nil
	}
}

func (c *Client) GetMetricData(ip, local, metricName, period string, start, stop int64, method int32) ([]model.Metric, error) {
	query := `from(bucket: "%s")
			  |> range(start: %v, stop: %v)
			  |> filter(fn: (r) => r["_measurement"] == "%s"
			  |> filter(fn: (r) => r["ip"] == "%s")
			  |> filter(fn: (r) => r["local"] == "%s")
			  |> filter(fn: (r) => r["_field"] == "value")
			  |> aggregateWindow(every: %s, fn: %s)`
	result, err := c.queryAPI.Query(context.Background(),
		fmt.Sprintf(query, c.bucket, start, stop+1, metricName, ip, local, period, Methods[method]))
	if err != nil {
		return nil, err
	}

	metrics := make([]model.Metric, 0, 20)
	for result.Next() {
		metrics = append(metrics, model.Metric{
			Name:      metricName,
			Value:     result.Record().Value().(float64),
			IP:        ip,
			Local:     local,
			Timestamp: result.Record().Time().Unix(),
		})
	}
	return metrics, nil
}

func parseTagsAndFileds(metric *model.Metric) (map[string]string, map[string]interface{}) {
	tags := make(map[string]string)
	tags["ip"] = metric.IP
	tags["local"] = metric.Local
	tags["port"] = metric.Port
	for k, v := range metric.Dimensions {
		tags[k] = v
	}
	return tags, map[string]interface{}{"value": metric.Value}
}
