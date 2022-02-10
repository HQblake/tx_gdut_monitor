package service

import (
	"context"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/cache"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/client"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/proto/judgpb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/worker"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/pkg/setting"
)

type Service struct {
	cache  *cache.Cache
	worker *worker.Worker
	send   send.ISend
	client *client.Client
	*judgpb.UnimplementedRuleUpdaterServer
}

func (s *Service) Update(ctx context.Context, rule *judgpb.AgentRule) (*judgpb.Response, error) {
	//TODO implement me
	panic("implement me")
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

	// 从协程池中提取worker完成各个指标的判定
	alert := s.worker.Finish(agent, rule, s.client)
	return s.send.Send(&alert)
}

func NewService(s *setting.Setting, send send.ISend) *Service {
	return &Service{
		cache:  cache.NewCache(s),
		worker: worker.NewWorker(s, judgment),
		send:   send,
		client: client.NewClient(s),
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
