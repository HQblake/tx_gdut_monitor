package influxdb

import (
	"bytes"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	"time"
)

func (c *Client) GetGetAllMethod() map[string]int32 {
	methods := make(map[string]int32)
	for k, v := range Methods {
		methods[v.Chinese] = k
	}
	return methods
}

// SaveMatricData 将上报记录保存到InfluxDB中
func (c *Client) SaveMatricData(metric *model.Metric) error {
	c.writeAPI.WriteRecord(parseRecord(metric))
	time.Now().Unix()
	return nil
}

func parseRecord(metric *model.Metric) string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%s, value=%v, local=%s, ip=%s, port=%s, ",
		metric.Name, metric.Value, metric.Local, metric.IP, metric.Port))
	for k, v := range metric.Dimensions {
		buf.WriteString(k + "=" + v + ", ")
	}
	buf.WriteString(string(metric.Timestamp))
	return buf.String()
}

func (c *Client) GetAggregatedData(name string, period string, method int32, timestamp int64) (float64, error) {
	return 0, nil
}

func (c *Client) GetMetricData(ip, local, metricName, period string, begin, end int64) []model.Metric {
	return nil
}
