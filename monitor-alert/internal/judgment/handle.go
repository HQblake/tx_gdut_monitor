package judgment

import (
	"bytes"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/proto"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"sync"
)

func judgment(metric string, agent *model.AgentReport, rule *model.AgentRule) (level int8) {
	// 对单个指标进行数据解析
	sql := parseMetricReport(metric, agent)

	// 从存储系统中获得聚合结果
	aggregation := proto.GetAggregation(sql, rule.Metrics[metric].Period, rule.Metrics[metric].Method)

	if threshold, ok := rule.Metrics[metric]; ok {
		for k, v := range threshold.Threshold {
			if k > level && v < aggregation {
				level = k
			}
		}
	}
	return
}

type task struct {
	metric string
	agent  *model.AgentReport
	rule   *model.AgentRule
	alert  *model.AlertInfo
	wg     *sync.WaitGroup
}

func handler(i interface{}) {
	args := i.(task)
	level := judgment(args.metric, args.agent, args.rule)
	args.alert.Metrics[args.metric] = model.MetricInfo{
		Metric:    args.metric,
		Value:     args.agent.Metrics[args.metric],
		Threshold: args.rule.Metrics[args.metric].Threshold[level],
		Method:    args.rule.Metrics[args.metric].Method,
		Level:     level,
		Duration:  args.rule.Metrics[args.metric].Period,
		Start:     string(args.agent.Timestamp),
	}
	args.wg.Done()
}

func parseMetricReport(metric string, agent *model.AgentReport) string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%s, value=%v, local=%s, ip=%s, port=%s, ",
		metric, agent.Metrics[metric], agent.Local, agent.IP, agent.Port))
	for k, v := range agent.Dimensions {
		buf.WriteString(k + "=" + v + ", ")
	}
	buf.WriteString(string(agent.Timestamp))
	return buf.String()
}
