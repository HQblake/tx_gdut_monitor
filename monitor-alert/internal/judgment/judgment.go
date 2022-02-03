package judgment

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/cache"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/proto"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/pkg/setting"
	"github.com/panjf2000/ants/v2"
	"log"
	"os"
	"sync"
	"time"
)

// JudgmentService 告警系统——判定服务模块
type JudgmentService struct {
}

var ruleCache *cache.Cache
var workers *ants.PoolWithFunc
var redisSetting = &setting.RedisSetting{}
var workerSetting = &setting.WorkersSetting{}

type task struct {
	metrix string
	agent  *model.AgentReport
	rule   *model.AgentRule
	alert  *model.AlertInfo
	wg     *sync.WaitGroup
}

// 为告警系统的其他服务提供调用 JudgmentService 服务相应功能的接口

func NewJudgmentService(s *setting.Setting) *JudgmentService {
	// 读取Redis配置
	s.ReadSection("Workers", workerSetting)
	s.ReadSection("Redis", redisSetting)

	ruleCache = cache.NewCache(redisSetting)
	duration, _ := time.ParseDuration(workerSetting.ExpiryDuration)

	var err error
	workers, err = ants.NewPoolWithFunc(workerSetting.Capacity, handler, ants.WithOptions(ants.Options{
		ExpiryDuration:   duration,
		PreAlloc:         workerSetting.PreAlloc,
		MaxBlockingTasks: workerSetting.MaxBlockingTasks,
		Nonblocking:      workerSetting.Nonblocking,
		PanicHandler:     nil,
		Logger:           ants.Logger(log.New(os.Stderr, "", log.LstdFlags)),
	}))
	if err != nil {
		log.Fatalln(err)
	}
	return &JudgmentService{}
}

// Check 方法用于判定接入服务接收的agent上报数据是否正常
func (js *JudgmentService) Check(agent *model.AgentReport) error {
	defer model.AgentReportPool.Put(agent)

	agentID := agent.Local + "-" + agent.IP

	// 获取该agent的判定规则
	rule := ruleCache.GetRuleByID(agentID)
	defer model.AgentRulePool.Put(rule)

	// 下面这段代码需要使用Ants协程池优化
	// 从协程池中提取workers完成各个指标的判定
	var wg sync.WaitGroup
	alert := model.AlertInfoPool.Get().(*model.AlertInfo)
	alert.Metrics = make(map[string]model.MetricInfo)
	alert.AgentID = agentID
	for k, _ := range agent.Metrics {
		wg.Add(1)
		_ = workers.Invoke(task{k, agent, rule, alert, &wg})
	}
	wg.Wait()

	// 将 alert 转发给发送服务

	return nil
}

// Update 方法接收管理服务发来的agent指标判定规则，更新服务内部缓存
func (js *JudgmentService) Update(ctx context.Context, req *proto.RuleReq) (*proto.RuleRsp, error) {
	rule := model.AgentRulePool.Get().(*model.AgentRule)
	defer model.AgentRulePool.Put(rule)

	err := json.Unmarshal([]byte(req.AgentRule), rule)
	if err != nil {
		return &proto.RuleRsp{Code: 500, Msg: err.Error()}, err
	}

	err = ruleCache.SetRuleByID(req.AgentID, rule)
	if err != nil {
		return &proto.RuleRsp{Code: 500, Msg: err.Error()}, err
	}
	return &proto.RuleRsp{Code: 200, Msg: "Success"}, nil
}

func judgment(metrix string, agent *model.AgentReport, rule *model.AgentRule) (level int8) {
	// 对单个指标进行数据解析
	sql := parseMetricReport(metrix, agent)

	// 从存储系统中获得聚合结果
	aggregation := proto.GetAggregation(sql, rule.Metrics[metrix].Period, rule.Metrics[metrix].Method)

	if threshold, ok := rule.Metrics[metrix]; ok {
		for k, v := range threshold.Threshold {
			if k > level && v < aggregation {
				level = k
			}
		}
	}
	return
}

func handler(i interface{}) {
	args := i.(task)
	level := judgment(args.metrix, args.agent, args.rule)
	args.alert.Metrics[args.metrix] = model.MetricInfo{
		Metric:    args.metrix,
		Value:     args.agent.Metrics[args.metrix],
		Threshold: args.rule.Metrics[args.metrix].Threshold[level],
		Method:    args.rule.Metrics[args.metrix].Method,
		Level:     level,
		Duration:  args.rule.Metrics[args.metrix].Period,
		Start:     string(args.agent.Timestamp),
	}
	args.wg.Done()
}

func parseMetricReport(metrix string, agent *model.AgentReport) string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%s, value=%v, local=%s, ip=%s, port=%s, ",
		metrix, agent.Metrics[metrix], agent.Local, agent.IP, agent.Port))
	for k, v := range agent.Dimensions {
		buf.WriteString(k + "=" + v + ", ")
	}
	buf.WriteString(string(agent.Timestamp))
	return buf.String()
}
