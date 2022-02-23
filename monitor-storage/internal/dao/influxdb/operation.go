package influxdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// SaveMatricData 将上报记录保存到InfluxDB中
func (c *Client) SaveMatricData(metric *model.Metric) error {
	tags, fields := parseTagsAndFileds(metric)
	c.writeAPI.WritePoint(influxdb2.NewPoint(metric.Name, tags, fields, time.Unix(metric.Timestamp, 0)))
	c.writeAPI.Flush()
	log.Printf("SaveMatricData: %v\n", *metric)
	return nil
}

func (c *Client) GetAggregatedData(metric *model.Metric, period string, method int32, timestamp int64) (float64, error) {
	query := `from(bucket: "%s")
	|> range(start: %v, stop: %v)
	|> filter(fn: (r) => r["_measurement"] == "%s")
	|> filter(fn: (r) => r["ip"] == "%s")
	|> filter(fn: (r) => r["local"] == "%s")
	|> filter(fn: (r) => r["_field"] == "value")
	|> aggregateWindow(every: %s, fn: %s, createEmpty: false)`
	duration, err := time.ParseDuration(period)
	if err != nil {
		return 0, err
	}
	stop := time.Unix(timestamp, 0).Add(1 * time.Second)
	start := stop.Add(-(duration + 1)).Unix()

	flux := fmt.Sprintf(query, c.bucket, start, stop.Unix(), metric.Name, metric.IP, metric.Local, period, Methods[method].English)
	result, err := c.queryAPI.Query(context.Background(), flux)
	if err != nil {
		return 0, err
	}

	var value float64
	if result.Next() {
		value = result.Record().Value().(float64)
	} else {
		value = metric.Value
	}

	log.Printf("GetAggregatedData(%v, %s, %s, %d): %f\n", *metric, period, Methods[method].English, timestamp, value)
	return value, nil
}

func (c *Client) GetMetricData(ip, local, metricName, period string, start, stop int64,
	method, limit int32) ([]model.Metric, error) {
	var query string
	if method < 0 {
		if limit <= 0 {
			query = fmt.Sprintf(everyWitoutLimit, c.bucket,
				start, stop+1, metricName, ip, local)
		} else {
			query = fmt.Sprintf(everyWitLimit, c.bucket,
				start, stop+1, metricName, ip, local, limit)
		}
	} else {
		m := Methods[method].English
		if limit <= 0 {
			query = fmt.Sprintf(aggregateWitoutLimit, c.bucket,
				start, stop+1, metricName, ip, local, period, m, false, m)
		} else {
			query = fmt.Sprintf(aggregateWitLimit, c.bucket,
				start, stop+1, metricName, ip, local, period, m, false, limit, m)
		}
	}
	fmt.Println(query)
	result, err := c.queryAPI.Query(context.Background(), query)
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

	log.Printf("GetMetricData(%s, %s, %s, %s, %s, %d, %d, %d): Found %d records\n",
		ip, local, metricName, period, Methods[method].English, start, stop, limit, len(metrics))
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
