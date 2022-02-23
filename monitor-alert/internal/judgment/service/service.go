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
	"runtime"
	"time"
)

type Service struct {
	cache  *cache.Cache
	pool   *ants.Pool
	send   send.ISend
	client *client.Client
	*judgpb.UnimplementedRuleUpdaterServer
}

type aggrate struct {
	metric string
	value  float64
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
	alert := model.AlertInfoPool.Get().(*model.AlertInfo)
	alert.IP, alert.Local, alert.Metrics = agent.IP, agent.Local, make(map[string]model.MetricInfo)

	// 创建聚合结果channel
	aggregations := make(chan aggrate, len(agent.Metrics))
	s.pool.Submit(func() {
		for k, _ := range agent.Metrics {
			aggregations <- aggrate{
				metric: k,
				value:  s.client.GetAggregation(k, agent, rule),
			}
		}
		close(aggregations)
	})

	for result := range aggregations {
		var level int32

		// 计算判定等级
		for k, v := range rule.Metrics[result.metric].Threshold {
			if k > level && v < result.value {
				level = k
			}
		}

		// 整合告警信息
		alert.Metrics[result.metric] = model.MetricInfo{
			Metric:    result.metric,
			Value:     result.value,
			Threshold: rule.Metrics[result.metric].Threshold[level],
			Method:    rule.Metrics[result.metric].Method,
			Level:     level,
			Duration:  rule.Metrics[result.metric].Period,
			Start:     agent.Timestamp,
		}

		// 保存告警信息
		s.client.SaveAlert(result.metric, alert)
	}
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
		PanicHandler: func(i interface{}) {
			defer func() {
				// 发生宕机时，获取panic传递的上下文并打印
				err := recover()
				switch err.(type) {
				case runtime.Error: // 运行时错误
					log.Println("runtime error:", err)
				default: // 非运行时错误
					log.Println("error:", err)
				}
			}()
		},
		Logger: ants.Logger(log.New(os.Stderr, "", log.LstdFlags)),
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
