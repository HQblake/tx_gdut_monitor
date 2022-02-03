package cache

import (
	"context"
	"encoding/json"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/proto"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/pkg/setting"
	"github.com/go-redis/redis/v8"
	"sync"
)

type Cache struct {
	mx  sync.RWMutex
	rdb *redis.Client
}

func NewCache(s *setting.RedisSetting) *Cache {
	return &Cache{rdb: redis.NewClient(&redis.Options{
		Addr:     s.Host,
		Password: s.Password,
		DB:       s.DB,
	})}
}

// GetRuleByID 需要优化
func (cache *Cache) GetRuleByID(agentID string) *model.AgentRule {
	cache.mx.Lock()
	rule := model.AgentRulePool.Get().(*model.AgentRule)
	rule.Metrics = make(map[string]model.MetricRule)

	// 从Redis中获取相应规则
	value, err := cache.rdb.Get(context.Background(), agentID).Result()
	if err == redis.Nil {
		// 缓存中无此规则，向管理服务请求规则
		value = proto.GetAgentRule(agentID)
		// 将此规则保存至缓存
		cache.setRuleByID(agentID, value)
	}
	cache.mx.Unlock()

	json.Unmarshal([]byte(value), rule)
	return rule
}

func (cache *Cache) SetRuleByID(id string, rule *model.AgentRule) error {
	cache.mx.Lock()
	defer cache.mx.Unlock()
	value, err := json.Marshal(rule)
	if err != nil {
		return err
	}
	return cache.setRuleByID(id, string(value))
}

func (cache *Cache) setRuleByID(id string, value string) error {
	return cache.rdb.Set(context.Background(), id, value, 0).Err()
}
