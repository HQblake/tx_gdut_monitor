package receive

import (
	"context"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/global"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/model"
	receivepb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/receive/proto"
	"google.golang.org/grpc/peer"
	"strings"
)

// ReceiveService 告警系统——判定服务模块
type ReceiveService struct {
	receivepb.UnimplementedReportServerServer
}

func (s *ReceiveService) Report(ctx context.Context, in *receivepb.ReportReq) (out *receivepb.ReportRsp, err error) {
	var agentreport model.AgentReport //封装AgentReport结构
	p, _ := peer.FromContext(ctx)     //获取Agent的IP地址和端口号 形式为IP:PORT
	//fmt.Println(p.Addr.String())  //打印Agent的IP地址和端口号
	str := strings.Split(p.Addr.String(), ":")
	agentreport.IP = str[0]                   //封装IP
	agentreport.Port = str[1]                 //封装port
	agentreport.Metrics = in.Metric           //封装指标Map
	agentreport.Timestamp = in.GetTimestamp() //封装时间戳
	agentreport.Local = in.Local              //封装区域
	//打印相应数据信息
	fmt.Println("Agent IP:", agentreport.IP)
	fmt.Println("Agent Port:", agentreport.Port)
	fmt.Println("Agent Location:", agentreport.Local)
	fmt.Println("TimeStamp:", agentreport.Timestamp)
	for k, str := range agentreport.Metrics {
		fmt.Println(k, ":", str)
	}
	judgment.NewService(global.Setting, global.Send).Check(&agentreport)
	//此区域之前已经封装好AgentReport结构，后面可写入判定模块的代码

	return &receivepb.ReportRsp{Code: receivepb.ResponseCode_SUCCESS, Msg: "成功返回"}, nil
}

//func main() { //此为服务端主函数
//	l, _ := net.Listen("tcp", ":8888")
//	s := grpc.NewServer()
//	receivepb.RegisterReportServerServer(s, &ReceiveService{})
//	s.Serve(l)
//}

// 为告警系统的其他服务提供调用 ReceiveService 服务相应功能的接口
