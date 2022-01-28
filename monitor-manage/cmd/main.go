package main

import (
	"flag"
	"log"
	"manage/configs"
	"manage/internal/service/admin"
)

// 监控系统的入口程序,也是主程序入口，在此做各个模块的初始化处理
func main() {
	flag.Parse()
	// 初始化配置
	err := configs.InitConfig(configPath)
	if err != nil {
		log.Fatalf("init config error %v", err)
	}


	// 下列代码仅作示例，可以参考，各个模块提供好自己的注册方式就行，也可以后续接口化统一，方便扩展
	// 注册管理模块，可以等完成http接口再打开注释
	_ = admin.Register()

	// 注册其他模块

}
