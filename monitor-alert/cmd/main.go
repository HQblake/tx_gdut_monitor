package main

import (
	"flag"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/global"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/receive"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/pkg/setting"
	"google.golang.org/grpc"
	"net"
)

func init() {
	var config string
	flag.StringVar(&config, "config", "configs/config.yaml", "告警系统配置文件")
	flag.Parse()

	setupSetting(config)
	setupService()
}

func main() {
	server := grpc.NewServer()
	// 注册发送服务
	global.Send.RegisterService(server)
	// 注册判定服务
	global.Judgement.RegisterService(server)
	// 注册接入服务
	global.Receive.RegisterService(server)

	// 启动端口监听
	lis, _ := net.Listen(global.Setting.Hosts.Network, global.Setting.Hosts.Server)
	server.Serve(lis)
}

func setupSetting(config string) {
	global.Setting, _ = setting.NewSetting(config)
}

func setupService() {
	// 服务加载顺序：先加载发送服务、再加载判定服务、最后加载接入服务
	global.Send = send.NewService()
	global.Judgement = judgment.NewService(global.Setting, global.Send)
	global.Receive = receive.NewService(global.Judgement)
	// 加载接入服务待实现
}
