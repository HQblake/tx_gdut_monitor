package service

import (
	"context"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/cache"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/client"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/proto/judgpb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/worker"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send"
	"log"
)

type Service struct {
	cache  *cache.Cache
	worker *worker.Worker
	send   send.ISend
	client *client.Client
	*judgpb.UnimplementedRuleUpdaterServer
}

func (s *Service) Update(ctx context.Context, rule *judgpb.AgentRule) (*judgpb.Response, error) {
	ruleTemp := model.AgentRulePool.Get().(*model.AgentRule)
	defer model.AgentRulePool.Put(ruleTemp)
	parseAgentRule(rule, ruleTemp)
	err := s.cache.SetRuleByID(rule.IP, rule.Local, ruleTemp)
	if err != nil {
		return &judgpb.Response{
			Code: judgpb.ResponseCode_ERROR,
			Msg:  err.Error(),
		}, err
	}
	return &judgpb.Response{
		Code: judgpb.ResponseCode_SUCCESS,
		Msg:  err.Error(),
	}, nil
}

func (s *Service) Check(agent *model.AgentReport) error {
	// 获取metric列表
	metrics := make([]string, 0, len(agent.Metrics))
	for k, _ := range agent.Metrics {
		metrics = append(metrics, k)
	}

	// 获取该agent的判定规则
	rule := model.AgentRulePool.Get().(*model.AgentRule)
	*rule = s.cache.GetRuleByIPAndLocal(agent.IP, agent.Local, metrics, s.client)
	defer model.AgentRulePool.Put(rule)
	log.Printf("Judgment rules: %v\n", *rule)

	// 从协程池中提取worker完成各个指标的判定
	alert := s.worker.Finish(agent, rule, s.client)
	log.Printf("Judgment result: %v\n", alert)

	return s.send.Send(&alert)
}

func NewService(send send.ISend) *Service {
	return &Service{
		cache:  cache.NewCache(),
		worker: worker.NewWorker(judgment),
		send:   send,
		client: client.NewClient(),
	}
}

func judgment(metric string, agent *model.AgentReport, rule *model.AgentRule, client *client.Client) (level int32) {
	// 从存储系统中获得聚合结果
	aggregation := client.GetAggregation(metric, agent, rule)
	if threshold, ok := rule.Metrics[metric]; ok {
		for k, v := range threshold.Threshold {
			if k > level && v < aggregation {
				level = k
			}
		}
	}
	return
}

func parseAgentRule(rule *judgpb.AgentRule, temp *model.AgentRule) {
	temp.IP = rule.IP
	temp.Local = rule.Local
	temp.Metrics = make(map[string]model.MetricRule)
	for k, v := range rule.Metrics {
		temp.Metrics[k] = model.MetricRule{
			Method:    v.Method,
			Period:    v.Period,
			Threshold: v.Threshold,
		}
	}
}
