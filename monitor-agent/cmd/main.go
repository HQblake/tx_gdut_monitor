package main

import (
	"context"
	"flag"
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-agent/internal/model"
	receivepb "gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-agent/internal/proto"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-agent/pkg/setting"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

var config string

func init() {
	flag.StringVar(&config, "config", "./configs/config.yaml", "Agent系统配置文件")
	flag.Parse()
}

func main() { //Agent端的主程序

	err := setting.InitConfig(config) //读取配置文件
	if err != nil {
		fmt.Println(err)
	}
	hostName := setting.GetConnectHost()   //从配置文件中获取服务端IP
	port := setting.GetConnectPort()       //获取服务端grpc监听的端口
	interval := setting.GetAgentInterval() //获取上报时间间隔 为了测试方便 配置文件中设置5
	metrics := setting.GetAgentMetrics()   //获取需要上传的指标
	local := setting.GetConnectLocation()  //获取Agent的区域
	Metrics := make(map[string]float64)    //准备好Map存储本机的指标值
	//for _, str := range metrics {
	//	f := global.GlobalMetrics[str]
	//	Metrics[str] = f()
	//	fmt.Println(f())
	//}
	coon, e := grpc.Dial(hostName+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials())) //准备发起grpc请求
	if e != nil {
		fmt.Println(e)
	}
	defer coon.Close()
	client := receivepb.NewReportServerClient(coon) //注册服务
	unixtime := time.Now().Unix()                   //获取当前时间戳
	for _, str := range metrics {
		f := model.GlobalMetrics[str]     //从全局Map中获取取得对应指标值的函数
		Metrics[str] = f() / float64(100) //将指标值填入Map
		fmt.Println(str, ":", Metrics[str])
	}
	//发起grpc请求，上报数据
	req, _ := client.Report(context.Background(), &receivepb.ReportReq{Timestamp: unixtime, Metric: Metrics, Local: local})
	fmt.Println(req.GetMsg())                                    //打印返回结果
	for range time.Tick(time.Duration(interval) * time.Second) { //定时器，以大小为Internal的固定时间间隔重复上报数据
		unixtime := time.Now().Unix() //获取当前时间戳
		for _, str := range metrics {
			f := model.GlobalMetrics[str]     //从全局Map中获取取得对应指标值的函数
			Metrics[str] = f() / float64(100) //将指标值填入Map
			fmt.Println(str, ":", Metrics[str])
		}
		//发起grpc请求，上报数据
		req, _ := client.Report(context.Background(), &receivepb.ReportReq{Timestamp: unixtime, Metric: Metrics, Local: local})
		fmt.Println(req.GetMsg()) //打印返回结果
	}
}
