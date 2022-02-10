package worker

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/client"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"sync"
)

type JudgFunc func(string, *model.AgentReport, *model.AgentRule, *client.Client) int32

type task struct {
	metric   string
	agent    *model.AgentReport
	rule     *model.AgentRule
	alert    *model.AlertInfo
	client   *client.Client
	wg       *sync.WaitGroup
	judgment JudgFunc
}

func handler(i interface{}) {
	args := i.(task)
	level := args.judgment(args.metric, args.agent, args.rule, args.client)
	args.alert.Metrics[args.metric] = model.MetricInfo{
		Metric:    args.metric,
		Value:     args.agent.Metrics[args.metric],
		Threshold: args.rule.Metrics[args.metric].Threshold[level],
		Method:    args.rule.Metrics[args.metric].Method,
		Level:     level,
		Duration:  args.rule.Metrics[args.metric].Period,
		Start:     args.agent.Timestamp,
	}
	// 将告警信息保存到存储系统中
	args.client.SaveAlert(args.metric, args.alert)
	args.wg.Done()
}
