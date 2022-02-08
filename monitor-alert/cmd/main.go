package main

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp",":8082")
	if err != nil {
		log.Fatal(err)
	}
	ser := grpc.NewServer()

	// 告警系统程序启动入口

	// 初始化接收模块

	// 初始化判定模块

	// 初始化发送模块,入参是供管理模块调用的grpc地址
	s := send.NewService()
	// 发送模块的功能接口（由判定模块看情况调用）
	//s.Send()
	// 注册rpc服务
	s.RegisterService(ser)
	err = ser.Serve(l)
	if err != nil {
		log.Println("grpc server:", err)
		ser.Stop()
		return
	}
	ser.Stop()



	// 暂时堵塞以便其他模块接入测试，之后将会是接收模块开启监听，堵塞主程序
	select {

	}

}
