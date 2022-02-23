package main

import (
	"flag"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/global/setting"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/receive"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	var config string
	flag.StringVar(&config, "config", "configs/config.yaml", "告警系统配置文件")
	flag.Parse()

	setupSetting(config)
	log.Println("配置文件加载完成")
	//setupService()
}

func main() {
	// 服务加载顺序：先加载发送服务、再加载判定服务、最后加载接入服务
	Send := send.NewService()
	log.Println("发送服务加载完成")
	Judgement := judgment.NewService(Send)
	log.Println("判定服务加载完成")
	Receive := receive.NewService(Judgement)
	log.Println("接入服务加载完成")

	server := grpc.NewServer(
		grpc.MaxConcurrentStreams(1000),
		grpc.NumStreamWorkers(1000))

	// 注册发送服务
	Send.RegisterService(server)
	// 注册判定服务
	Judgement.RegisterService(server)
	// 注册接入服务
	Receive.RegisterService(server)

	// 启动端口监听
	host := setting.GetHostConfig()
	lis, _ := net.Listen(host.Network, host.Server)
	server.Serve(lis)
}

func setupSetting(config string) {
	err := setting.InitSetting(config)
	if err != nil {
		log.Fatal(err)
	}
}

//
//func setupService() {
//	// 服务加载顺序：先加载发送服务、再加载判定服务、最后加载接入服务
//	Send := send.NewService()
//	log.Println("发送服务加载完成")
//	Judgement := judgment.NewService(Send)
//	log.Println("判定服务加载完成")
//	Receive := receive.NewService(Judgement)
//	log.Println("接入服务加载完成")
//}
