package client

import (
	"context"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/proto/managepb"
	storepb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment/proto/storagepb"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/pkg/setting"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client struct {
	manage  managepb.ManageServiceClient
	storage storepb.MetricServiceClient
}

func (c *Client) GetAggregation(metric string, agent *model.AgentReport, rule *model.AgentRule) float64 {
	resp, err := c.storage.GetAggregatedData(context.Background(), &storepb.AggregatedRequest{
		Method:     int32(rule.Metrics[metric].Method),
		Period:     rule.Metrics[metric].Period,
		Metric:     metric,
		Value:      agent.Metrics[metric],
		IP:         agent.IP,
		Local:      agent.Local,
		Port:       agent.Port,
		Timestamp:  agent.Timestamp,
		Dimensions: agent.Dimensions,
	})
	if err != nil {
		log.Println(err)
	}
	if resp.Code == storepb.ResponseCode_ERROR {
		log.Println(err)
	}
	return resp.Result
}

func (c *Client) SaveAlert(metric string, alert *model.AlertInfo) {
	resp, err := c.storage.InsertAlertInfo(context.Background(), &storepb.HistoryInfoRequest{
		IP:        alert.IP,
		Local:     alert.Local,
		Metric:    metric,
		Value:     alert.Metrics[metric].Value,
		Threshold: alert.Metrics[metric].Threshold,
		Method:    alert.Metrics[metric].Method,
		Level:     alert.Metrics[metric].Level,
		Duration:  alert.Metrics[metric].Duration,
		Start:     alert.Metrics[metric].Start,
	})
	if err != nil {
		log.Println(err)
	}
	if resp.Code == storepb.ResponseCode_ERROR {
		log.Println(resp.Msg)
	}
}

func (c *Client) GetAgentRule(ip, local string, metrics []string) model.AgentRule {
	resp, err := c.manage.Get(context.Background(), &managepb.CheckRequest{
		IP:      ip,
		Local:   local,
		Metrics: metrics,
	})
	if err != nil {
		log.Println(err)
	}

	agent := model.AgentRule{}
	if resp.Code == managepb.ResponseCode_ERROR {
		log.Println(resp.Msg)
	} else {
		agent.IP = resp.Result.IP
		agent.Local = resp.Result.Local
		agent.Metrics = make(map[string]model.MetricRule)
		for k, v := range resp.Result.Metrics {
			agent.Metrics[k] = model.MetricRule{
				Method:    v.Method,
				Period:    v.Period,
				Threshold: v.Threshold,
			}
		}
	}
	return agent
}

func NewClient(s *setting.Setting) *Client {
	client := &Client{}
	var conn *grpc.ClientConn
	var err error

	conn, err = grpc.Dial(s.Hosts.ManageClient, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	client.manage = managepb.NewManageServiceClient(conn)

	conn, err = grpc.Dial(s.Hosts.StorageClient, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	client.storage = storepb.NewMetricServiceClient(conn)
	return client
}
