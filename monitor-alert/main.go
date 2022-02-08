package main

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/global"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/judgment"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/pkg/setting"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	setupSetting()
	setupService()
}

func main() {

}

func setupSetting() {
	global.Setting, _ = setting.NewSetting()
}

func setupService() {
	l, err := net.Listen("tcp",":8082")
	if err != nil {
		log.Fatal(err)
	}
	ser := grpc.NewServer()
	// 加载判定服务

	global.JudgementService, err = judgment.NewJudgmentService(global.Setting)
	if err != nil {
		log.Fatalln(err)
	}

	// 加载发送服务
	global.SendService, err = send.NewService()
	if err != nil {
		log.Fatalln(err)
	}
	// 注册发送服务的rpc
	global.SendService.RegisterService(ser)
	err = ser.Serve(l)
	if err != nil {
		log.Println("grpc server:", err)
		ser.Stop()
		return
	}
	ser.Stop()
}
