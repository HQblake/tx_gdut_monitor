package main

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/global"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/global/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/setting"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	var err error
	global.Setting, err = setting.NewSetting()
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
	lis, _ := net.Listen("tcp", ":")
	server.Serve(lis)
}
