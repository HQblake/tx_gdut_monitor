package judgment

import (
	"context"
	"encoding/json"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/global"
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
type JudgmentService struct{}

var ruleCache *cache.Cache
var workers *ants.PoolWithFunc
var redisSetting = &setting.RedisSetting{}
var workerSetting = &setting.WorkersSetting{}

// 为告警系统的其他服务提供调用 JudgmentService 服务相应功能的接口

// NewJudgmentService 创建新的判定服务服务
func NewJudgmentService(s *setting.Setting) (*JudgmentService, error) {
	// 读取协程池与Redis的配置
	err := s.ReadSection("Workers", workerSetting)
	if err == nil {
		return nil, err
	}

	s.ReadSection("Redis", redisSetting)
	if err == nil {
		return nil, err
	}

	ruleCache = cache.NewCache(redisSetting)
	duration, _ := time.ParseDuration(workerSetting.ExpiryDuration)

	workers, err = ants.NewPoolWithFunc(workerSetting.Capacity, handler, ants.WithOptions(ants.Options{
		ExpiryDuration:   duration,
		PreAlloc:         workerSetting.PreAlloc,
		MaxBlockingTasks: workerSetting.MaxBlockingTasks,
		Nonblocking:      workerSetting.Nonblocking,
		PanicHandler:     nil,
		Logger:           ants.Logger(log.New(os.Stderr, "", log.LstdFlags)),
	}))
	if err != nil {
		return nil, err
	}
	return &JudgmentService{}, nil
}

// Check 方法用于判定接入服务接收的agent上报数据是否正常
func (js *JudgmentService) Check(agent *model.AgentReport) error {
	defer model.AgentReportPool.Put(agent)

	agentID := agent.IP + " " + agent.Local

	// 获取该agent的判定规则
	rule := ruleCache.GetRuleByID(agentID)
	defer model.AgentRulePool.Put(rule)

	// 从协程池中提取workers完成各个指标的判定
	var wg sync.WaitGroup
	alert := model.AlertInfoPool.Get().(*model.AlertInfo)
	defer model.AlertInfoPool.Put(alert)
	alert.Metrics = make(map[string]model.MetricInfo)
	alert.AgentID = agentID
	for k, _ := range agent.Metrics {
		wg.Add(1)
		_ = workers.Invoke(task{k, agent, rule, alert, &wg})
	}
	wg.Wait()

	// 将 alert 转发给发送服务
	return (*global.SendService).Send(alert)
}

// Update 方法接收管理服务发来的agent指标判定规则，更新服务内部缓存
func (js *JudgmentService) Update(ctx context.Context, req *proto.RuleReq) (*proto.RuleRsp, error) {
	rule := model.AgentRulePool.Get().(*model.AgentRule)
	rule.Metrics = make(map[string]model.MetricRule)
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
