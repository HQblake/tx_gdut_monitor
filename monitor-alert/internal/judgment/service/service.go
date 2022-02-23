package service

import (
	"context"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/global/setting"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/cache"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/client"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/proto/judgpb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send"
	"github.com/panjf2000/ants/v2"
	"log"
	"os"
	"sync"
	"time"
)

type Service struct {
	cache  *cache.Cache
	pool   *ants.Pool
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

	// 获取从缓存中获取该agent的判定规则
	rule := model.AgentRulePool.Get().(*model.AgentRule)
	*rule = s.cache.GetRuleByIPAndLocal(agent.IP, agent.Local, metrics, s.client)
	defer model.AgentRulePool.Put(rule)
	log.Printf("Judgment rules: %v\n", *rule)

	// 从协程池中提取worker完成各个指标的判定
	var wg sync.WaitGroup
	alert := model.AlertInfoPool.Get().(*model.AlertInfo)
	alert.IP, alert.Local, alert.Metrics = agent.IP, agent.Local, make(map[string]model.MetricInfo)
	for k, _ := range agent.Metrics {
		wg.Add(1)
		s.pool.Submit(func() {
			var level int32

			// 从存储系统中获得聚合结果
			aggregation := s.client.GetAggregation(k, agent, rule)
			if threshold, ok := rule.Metrics[k]; ok {
				for k, v := range threshold.Threshold {
					if k > level && v < aggregation {
						level = k
					}
				}
			}

			// 整合告警信息
			alert.Metrics[k] = model.MetricInfo{
				Metric:    k,
				Value:     agent.Metrics[k],
				Threshold: rule.Metrics[k].Threshold[level],
				Method:    rule.Metrics[k].Method,
				Level:     level,
				Duration:  rule.Metrics[k].Period,
				Start:     agent.Timestamp,
			}

			// 将告警信息保存到存储系统中
			s.client.SaveAlert(k, alert)
			wg.Done()
		})
	}
	wg.Wait()
	log.Printf("Judgment result: %v\n", alert)

	// 发送告警信息
	s.pool.Submit(func() {
		s.send.Send(alert)
		model.AlertInfoPool.Put(alert)
	})

	return nil
}

func NewService(send send.ISend) *Service {
	// 初始化协程池
	ws, err := setting.GetWorkerConfig()
	if err != nil {
		log.Fatalln(err)
	}
	duration, _ := time.ParseDuration(ws.ExpiryDuration)
	pool, err := ants.NewPool(ws.Capacity, ants.WithOptions(ants.Options{
		ExpiryDuration:   duration,
		PreAlloc:         ws.PreAlloc,
		MaxBlockingTasks: ws.MaxBlockingTasks,
		Nonblocking:      ws.Nonblocking,
		PanicHandler:     nil,
		Logger:           ants.Logger(log.New(os.Stderr, "", log.LstdFlags)),
	}))
	if err != nil {
		log.Fatalln(err)
	}

	return &Service{
		cache:  cache.NewCache(),
		pool:   pool,
		send:   send,
		client: client.NewClient(),
	}
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
