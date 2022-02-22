package cache

import (
	"context"
	"encoding/json"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/global/setting"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/client"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"github.com/go-redis/redis/v8"
	"log"
)

type Cache struct {
	mx  *CacheLock
	rdb *redis.Client
}

func NewCache() *Cache {
	//rs := RedisSetting{}
	//s.ReadSection("Redis", &rs)
	rs, err := setting.GetRedisConfig()
	if err != nil {
		log.Fatalln(err)
	}
	return &Cache{rdb: redis.NewClient(&redis.Options{
		Addr:     rs.Host,
		Password: rs.Password,
		DB:       rs.DB,
	}), mx: NewLock()}
}

// GetRuleByIPAndLocal 根据IP与Local获取规则
func (cache *Cache) GetRuleByIPAndLocal(ip, local string, metrics []string, client *client.Client) model.AgentRule {
	agentID := local + "-" + ip
	cache.mx.Lock(agentID)
	defer cache.mx.Unlock(agentID)

	// 从Redis中获取相应规则
	value, err := cache.rdb.Get(context.Background(), agentID).Result()
	if err == redis.Nil {
		// 缓存中无此规则，向管理服务请求规则
		rule := client.GetAgentRule(ip, local, metrics)

		// 将此规则保存至缓存
		buf, _ := json.Marshal(&rule)
		cache.setRuleByID(agentID, string(buf))
		return rule
	}

	// 解析JSON字符，判定指标是否一致
	rule := model.AgentRulePool.Get().(*model.AgentRule)
	defer model.AgentRulePool.Put(rule)
	_ = json.Unmarshal([]byte(value), rule)

	n := len(rule.Metrics)
	newMetrics := make([]string, 0, n)
	for k, _ := range rule.Metrics {
		newMetrics = append(newMetrics, k)
	}
	for _, v := range metrics {
		if _, ok := rule.Metrics[v]; !ok {
			newMetrics = append(newMetrics, v)
		}
	}
	if len(newMetrics) != n {
		// 缓存规则过期，向管理服务请求新的规则
		*rule = client.GetAgentRule(rule.IP, rule.Local, newMetrics)
		// 将此规则保存至缓存
		buf, _ := json.Marshal(rule)
		cache.setRuleByID(agentID, string(buf))
	}
	return *rule
}

func (cache *Cache) SetRuleByID(ip, local string, rule *model.AgentRule) error {
	agentID := local + "-" + ip
	cache.mx.Lock(agentID)
	defer cache.mx.Unlock(agentID)
	value, err := json.Marshal(rule)
	if err != nil {
		return err
	}
	return cache.setRuleByID(agentID, string(value))
}

func (cache *Cache) setRuleByID(id string, value string) error {
	return cache.rdb.Set(context.Background(), id, value, 0).Err()
}
