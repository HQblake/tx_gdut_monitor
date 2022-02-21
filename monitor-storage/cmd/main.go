package main

import (
	"flag"
	"log"
	"net"

	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/global"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/global/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
	"google.golang.org/grpc"
)

func init() {
	var config string
	flag.StringVar(&config, "config", "./configs/config.yaml", "存储系统配置文件")
	flag.Parse()

	var err error
	global.Setting, err = setting.NewSetting(config)
	if err != nil {
		log.Fatalln(err)
	}
	global.JudgmentService = service.NewJudgmentService(global.Setting)
	global.ManageService = service.NewManageService(global.Setting)
}

func main() {
	server := grpc.NewServer()
	global.JudgmentService.RegisterJudgmentService(server)
	global.ManageService.RegisterManageService(server)
	lis, _ := net.Listen(global.Setting.Hosts.Network, global.Setting.Hosts.Server)
	server.Serve(lis)
}
