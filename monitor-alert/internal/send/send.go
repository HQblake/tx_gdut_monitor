package send

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	sendpb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/api/gen"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/api/service"
	model2 "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send/output"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

// ISend 接入判定服务
type ISend interface {
	// Send 对判定服务发来的告警信息进行处理，进行通知
	Send(alert *model.AlertInfo) error
	// RegisterService 注册rpc服务
	RegisterService(ser *grpc.Server)
}

type Service struct {
	proxy    *service.Service
	agents   output.IManager
	infoPool *sync.Pool
}



// NewService 初始化发送服务，提供对外的判定服务结构（判定服务直接调用该结构的Send方法即可完成发送）
func NewService() *Service{
	Register()
	agents := output.NewManager()
	//agents.Agents["127.0.0.1-test"] = output.NewOutputs("127.0.0.1-test")
	//err := agents.GetOutputs("127.0.0.1-test").Set(1, output.Config{Name: "email", Level: 0, Config: `{"target":"526756656@qq.com", "format_type":"html"}`})
	//if err != nil {
	//	log.Println(err)
	//}
	return &Service{
		agents: agents,
		proxy:  service.NewService(agents),
		infoPool: &sync.Pool{
			New: func() interface{} {
				return model2.Info{}
			},
		},
	}
}
func (s *Service) RegisterService(ser *grpc.Server) {
	sendpb.RegisterSendServiceServer(ser, s.proxy)
}

func (s *Service) Send(alert *model.AlertInfo) error {
	fmt.Printf("开始告警：%+v",alert)
	outputs := s.agents.GetOutputs(fmt.Sprintf("%s-%s", alert.IP, alert.Local))
	for _, info := range alert.Metrics {
		// 考虑开协程去分别处理,同一时间上报的指标数不会过多，暂不做协程数量限制
		i := s.newInfo(fmt.Sprintf("%s-%s", alert.IP, alert.Local), info)
		err := outputs.Output(i)
		if err != nil {
			log.Println(err)
		}
		s.release(i)
		//go func(i model2.Info) {
		//	err := outputs.Output(i)
		//	if err != nil {
		//		log.Println(err)
		//	}
		//	s.release(i)
		//}(i)
	}
	return nil
}

func (s *Service) newInfo(agent string, alert model.MetricInfo) model2.Info {
	i := s.infoPool.Get().(model2.Info)
	i.Agent = agent
	i.Metric = alert.Metric
	i.Value = alert.Value
	i.Threshold = alert.Threshold
	i.Level = output.Level(alert.Level).String()
	i.Duration = alert.Duration
	i.Start = time.Unix(alert.Start, 0).Format("[2006-01-01 15:04:05]")
	i.ParseMethod(alert.Method)
	return i
}

func (s *Service) release(info model2.Info) {
	s.infoPool.Put(info)
}
