package main

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-alert/internal/send"
	"log"
)

func main() {
	// 告警系统程序启动入口

	// 初始化接收模块

	// 初始化判定模块

	// 初始化发送模块,入参是供管理模块调用的grpc地址
	s, err := send.NewService(":8082")
	if err != nil{
		log.Fatal(err)
	}
	// 发送模块的功能接口（由判定模块看情况调用）
	//s.Send()

	fmt.Println(s)



	// 暂时堵塞，之后将会是接收模块开启监听，堵塞主程序
	select {

	}

}
