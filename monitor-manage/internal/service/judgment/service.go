package judgment

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/configs"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"
	managepb2 "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/judgment/gen"
	managepb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/store/gen"
	"io"
	"log"
)

type Service struct {
	store managepb.JudgmentServiceClient
	judgment managepb2.RuleUpdaterClient
}

func NewService(store managepb.JudgmentServiceClient, judgment managepb2.RuleUpdaterClient) *Service {
	return &Service{
		store: store,
		judgment: judgment,
	}
}

func (s *Service) GetConfigs(ip string, local string) (map[string]model.JudgmentConfig, error) {
	var err error
	// 获取存储服务中对应agent的所有判定规则
	stream, err := s.store.GetConfigsByAgent(context.Background(), &managepb.AgentRequest{IP: ip, Local: local})
	if err != nil {
		return nil, err
	}
	var resp *managepb.JudgmentConfigResponse
	res := make(map[string]model.JudgmentConfig)
	// 遍历获取指定metric的规则，没有则用默认规则代替
	for {
		resp, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("grpc get judgment rule error %v", err)
			continue
		}
		if resp.Code != managepb.ResponseCode_SUCCESS{
			log.Printf("grpc get judgment rule error %v", resp.Msg)
			continue
		}
		config := resp.GetResult()
		var threshold map[int32]float64
		err = json.Unmarshal([]byte(config.Threshold), &threshold)
		if err != nil {
			log.Printf("grpc get judgment json parse threshold rule error %v", err)
			continue
		}
		conf := model.JudgmentConfig{
			ID: config.GetID(),
			IP: config.GetIP(),
			Local: config.GetLocal(),
			Metric: config.GetMetric(),
			Method: config.GetMethod(),
			Period: config.GetPeriod(),
			Threshold: threshold,
		}
		res[config.GetMetric()] = conf
	}
	return res, nil
}

func (s *Service) GetConfigsWithMetrics(ip string, local string, metrics []string) ([]model.JudgmentConfig, error) {
	res := make([]model.JudgmentConfig, 0, len(metrics))
	cfgs, err := s.GetConfigs(ip, local)
	if err != nil {
		return nil, err
	}
	defaultRule := configs.GetDefaultRule()
	for _, m := range metrics {
		if _, ok := cfgs[m]; ok {
			res = append(res, cfgs[m])
		}else {
			res = append(res, model.JudgmentConfig{
				ID: -1,
				IP: ip,
				Local: local,
				Metric: m,
				Method: defaultRule.Method,
				Period: defaultRule.Period,
				Threshold: defaultRule.Threshold,
			})
		}
	}
	return res, nil
}

func (s *Service) Update(ID int32, IP string,Local string, Metric string, Method int32, Period string, Threshold string) error {
	resp, err := s.store.UpdateConfig(context.Background(), &managepb.JudgmentEntry{
		ID: ID,
		IP: IP,
		Local: Local,
		Metric: Metric,
		Method: Method,
		Period: Period,
		Threshold: Threshold,
	})
	if err != nil {
		return err
	}
	if resp.GetCode() != managepb.ResponseCode_SUCCESS {
		return fmt.Errorf("update judgment config store service error %s", resp.GetMsg())
	}
	go func(ip string, local string) {
		err := s.TriggerUpdate(ip, local)
		if err != nil {
			log.Printf("update judgment config trigger update error: %s", err.Error())
		}
	}(IP, Local)
	return nil
}

func (s *Service) Del(ip string, local string, id int32) error {
	resp, err := s.store.DeleteConfig(context.Background(), &managepb.IDRequest{ID: id})
	if err != nil {
		return err
	}
	if resp.GetCode() != managepb.ResponseCode_SUCCESS {
		return fmt.Errorf("del judgment config store service error %s", resp.GetMsg())
	}
	go func(ip string, local string) {
		err := s.TriggerUpdate(ip, local)
		if err != nil {
			log.Printf("del judgment config trigger update error: %s", err.Error())
		}
	}(ip, local)
	return nil
}

func (s *Service) TriggerUpdate(ip string, local string) error {
	list, err := s.GetConfigs(ip, local)
	if err != nil {
		return err
	}
	agent := &managepb2.AgentRule{
		IP: ip,
		Local: local,
		Metrics: make(map[string]*managepb2.MetricRule),
	}
	for key, config := range list {
		agent.Metrics[key] = &managepb2.MetricRule{
			Method: config.Method,
			Period: config.Period,
			Threshold: config.Threshold,
		}
	}
	resp, err := s.judgment.Update(context.Background(), agent)
	if err != nil {
		return err
	}
	if resp.GetCode() != managepb2.ResponseCode_SUCCESS {
		return fmt.Errorf("update judgment rpc service error %s", resp.GetMsg())
	}
	return nil
}


