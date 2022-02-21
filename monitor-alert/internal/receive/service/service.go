package service

import (
	"context"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/receive/receivepb"
	"google.golang.org/grpc/peer"
	"log"
	"strings"
)

// ReceiveService 告警系统——判定服务模块
type ReceiveService struct {
	judgment judgment.IJudgment
	receivepb.UnimplementedReportServerServer
}

func (s *ReceiveService) Report(ctx context.Context, in *receivepb.ReportReq) (*receivepb.ReportRsp, error) {
	agentreport := model.AgentReportPool.Get().(*model.AgentReport)
	defer model.AgentReportPool.Put(agentreport)
	p, _ := peer.FromContext(ctx) //获取Agent的IP地址和端口号 形式为IP:PORT
	//fmt.Println(p.Addr.String())  //打印Agent的IP地址和端口号
	str := strings.Split(p.Addr.String(), ":")

	// 用于压力测试的代码，后续可删除
	ip, port := str[0], str[1]
	if in.IP != "" {
		ip = in.IP
	}
	if in.Port != "" {
		port = in.Port
	}

	agentreport.IP = ip                       //封装IP
	agentreport.Port = port                   //封装port
	agentreport.Metrics = in.Metric           //封装指标Map
	agentreport.Timestamp = in.GetTimestamp() //封装时间戳
	agentreport.Local = in.Local              //封装区域
	for k, str := range agentreport.Metrics {
		fmt.Println(k, ":", str)
	}
	log.Printf("Agent reports data: %v\n", *agentreport)

	//此区域之前已经封装好AgentReport结构，后面可写入判定模块的代码
	s.judgment.Check(agentreport)

	return &receivepb.ReportRsp{Code: receivepb.ResponseCode_SUCCESS, Msg: "成功返回"}, nil
}

func NewService(judgment judgment.IJudgment) *ReceiveService {
	return &ReceiveService{judgment: judgment}
}

//func main() { //此为服务端主函数
//	l, _ := net.Listen("tcp", ":8888")
//	s := grpc.NewServer()
//	receivepb.RegisterReportServerServer(s, &ReceiveService{})
//	s.Serve(l)
//}

// 为告警系统的其他服务提供调用 ReceiveService 服务相应功能的接口
