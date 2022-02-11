package send

import (
	"context"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/model"
	sendpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/send/gen"
	managepb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/rpc/client/store/gen"
	"io"
	"log"
	"strconv"
	"sync"
)

type Service struct {
	once sync.Once
	send sendpb.SendServiceClient
	store managepb.SendServiceClient
}

func NewService(send sendpb.SendServiceClient, store managepb.SendServiceClient) *Service {
	return &Service{
		once: sync.Once{},
		send: send,
		store: store,
	}
}

func (s *Service) GetAllConfigs() ([]model.SendConfig, error) {
	var err error
	// 获取存储服务中对应agent的所有判定规则
	stream, err := s.store.GetAllConfigs(context.Background(), &managepb.BaseRequest{})
	if err != nil {
		return nil,  err
	}
	var resp *managepb.SendConfigResponse
	res := make([]model.SendConfig, 0, 10)
	// 遍历获取指定metric的规则，没有则用默认规则代替
	for {
		resp, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("grpc get all send config error %v", err)
			continue
		}
		if resp.Code != managepb.ResponseCode_SUCCESS{
			log.Printf("grpc get all send config error %v", resp.Msg)
			continue
		}
		config := resp.GetResult()
		res = append(res, model.SendConfig{
			ID: config.GetID(),
			IP: config.GetIP(),
			Local: config.GetLocal(),
			SendType:config.GetSendType(),
			Level:config.GetLevel(),
			Config:config.GetConfig(),
		})
	}
	return res, nil
}

func (s *Service) GetConfigs(ip string, local string) ([]model.SendConfig, error) {
	var err error
	// 获取存储服务中对应agent的所有判定规则
	stream, err := s.store.GetConfigsByAgent(context.Background(), &managepb.AgentRequest{IP: ip, Local: local})
	if err != nil {
		return nil, err
	}
	var resp *managepb.SendConfigResponse
	res := make([]model.SendConfig, 0)
	// 遍历获取指定metric的规则，没有则用默认规则代替
	for {
		resp, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("grpc get send config error %v", err)
			continue
		}
		if resp.Code != managepb.ResponseCode_SUCCESS{
			log.Printf("grpc get send config error %v", resp.Msg)
			continue
		}
		config := resp.GetResult()
		res = append(res, model.SendConfig{
			ID: config.GetID(),
			IP: config.GetIP(),
			Local: config.GetLocal(),
			SendType:config.GetSendType(),
			Level:config.GetLevel(),
			Config:config.GetConfig(),
		})
	}
	return res, nil
}

func (s *Service) AddConfig(Ip string, Local string, SendType int32, Level int32, Config string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := s.store.AddConfig(ctx, &managepb.AddSendRequest{
		IP: Ip,
		Local: Local,
		SendType: SendType,
		Level: Level,
		Config: Config,
	})
	if err != nil {
		return err
	}
	if resp.GetCode() != managepb.ResponseCode_SUCCESS {
		return fmt.Errorf("add send config store service error %s", resp.GetMsg())
	}
	id, err := strconv.Atoi(resp.GetMsg())
	if err != nil {
		return err
	}
	res, err := s.send.Set(ctx, &sendpb.UpdateRequest{
		IP: Ip,
		Local: Local,
		Config: &sendpb.ConfigEntry{
			ConfigID: int32(id),
			Conf: &sendpb.Config{
				SendType: sendpb.Type(SendType),
				Config: Config,
				Level: Level,
			},
		},
	})
	if err != nil {
		return err
	}
	if res.GetCode() != sendpb.SendResponse_SUCCESS {
		return fmt.Errorf("set send config send service error %s", res.GetMsg())
	}
	return nil
}

func (s *Service) Update(id int32, IP string, Local string, SendType int32, Level int32, Config string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// set包括检查，先set
	res, err := s.send.Set(ctx, &sendpb.UpdateRequest{
		IP: IP,
		Local: Local,
		Config: &sendpb.ConfigEntry{
			ConfigID: id,
			Conf: &sendpb.Config{
				SendType: sendpb.Type(SendType),
				Config: Config,
				Level: Level,
			},
		},
	})
	if err != nil {
		return err
	}
	if res.GetCode() != sendpb.SendResponse_SUCCESS {
		return fmt.Errorf("set send config send service error %s", res.GetMsg())
	}
	resp, err := s.store.UpdateConfig(ctx, &managepb.SendEntry{
		ID: id,
		IP: IP,
		Local: Local,
		SendType: SendType,
		Level: Level,
		Config: Config,
	})
	if err != nil {
		return err
	}
	if resp.GetCode() != managepb.ResponseCode_SUCCESS {
		return fmt.Errorf("update send config store service error %s", resp.GetMsg())
	}
	return nil
}

func (s *Service) Del(ip string, local string, id int32) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := s.store.DeleteConfig(ctx, &managepb.IDRequest{
		ID: id,
	})
	if err != nil {
		return err
	}
	if resp.GetCode() != managepb.ResponseCode_SUCCESS {
		return fmt.Errorf("del send config store service error %s", resp.GetMsg())
	}
	res, err := s.send.Del(ctx, &sendpb.DelRequest{
		IP: ip,
		Local: local,
		ConfigID: id,
	})
	if err != nil {
		return err
	}
	if res.GetCode() != sendpb.SendResponse_SUCCESS {
		return fmt.Errorf("del send config send service error %s", res.GetMsg())
	}
	return nil
}

func (s *Service) Init() {
	s.once.Do(func() {
		list, err := s.GetAllConfigs()
		if err != nil {
			log.Println(err)
		}
		res := make(map[string]*sendpb.InitConfig)
		for _, config := range list {
			agent := fmt.Sprintf("%s-%s", config.IP, config.Local)
			entry :=  &sendpb.ConfigEntry{
				ConfigID: config.ID,
				Conf: &sendpb.Config{
					SendType: sendpb.Type(config.SendType),
					Config: config.Config,
					Level: config.Level,
				},
			}
			if v, ok := res[agent]; ok {
				v.Config = append(v.Config,entry)
			}else {
				res[agent] = &sendpb.InitConfig{
					IP: config.IP,
					Local: config.Local,
					Config: make([]*sendpb.ConfigEntry, 0, 4),
				}
				res[agent].Config = append(res[agent].Config, entry)
			}
		}
		resp, err := s.send.Init(context.Background(), &sendpb.InitRequest{
			Config: res,
		})
		if err != nil {
			log.Println(err)
		}
		if resp.GetCode() != sendpb.SendResponse_SUCCESS {
			log.Printf("init send service error %s", resp.GetMsg())
		}
	})
}



