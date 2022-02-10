package main

import (
	"flag"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/configs"
	"log"

)
var configPath string

func init() {
	flag.StringVar(&configPath, "config", "./config.yml", "config path")
}
// 监控系统的入口程序,也是主程序入口，在此做各个模块的初始化处理
func main() {
	flag.Parse()
	// 初始化配置
	err := configs.InitConfig(configPath)
	if err != nil {
		log.Fatalf("init config error %v", err)
	}


}
