package main

import (
	"flag"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/configs"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)
var configPath string
var debug bool

func init() {
	flag.StringVar(&configPath, "config", "./config.yml", "config path")
	flag.BoolVar(&debug, "debug", false, "is debug")
}
// 监控系统的入口程序,也是主程序入口，在此做各个模块的初始化处理
func main() {
	flag.Parse()
	// 初始化配置
	err := configs.InitConfig(configPath)
	if err != nil {
		log.Fatalf("init config error %v", err)
	}
	// 监控系统的rpc地址
	alertConn, err := grpc.Dial(configs.GetAlertConnAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer alertConn.Close()
	// 存储系统的rpc地址
	storeConn, err := grpc.Dial(configs.GetStoreConnAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer storeConn.Close()

	admin := service.Register(alertConn, storeConn)
	//// 管理服务开启grpc通信
	//judgpb.RegisterManageServiceServer(a.Service)
	// 开启监听
	log.Fatal(router(admin.ApiHandler).Run(configs.GetAdminListenAddr()))

}

func router(h *http.Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	// agent
	agent := r.Group("/agent")
	agent.GET("/list", h.GetAllAgent)
	agent.GET("/info/:ip/:local", h.GetAgentInfo)

	// judgment
	judgment := r.Group("/judgment")
	judgment.GET("/info/:ip/:local", h.GetAllRule)
	judgment.POST("/update/:ip/:local/:id", h.UpdateRule)
	judgment.POST("/del/:ip/:local/:id", h.DelRule)

	// send
	send := r.Group("send")
	send.GET("/info/:ip/:local", h.GetSendConfigs)
	send.POST("/add/:ip/:local", h.AddSendConfig)
	send.POST("/update/:ip/:local/:id", h.UpdateSendConfig)
	send.POST("/del/:ip/:local/:id", h.DelSendConfig)


	// show
	show := r.Group("show")
	show.GET("/metrics", h.GetMetricsInOneDay)

	return r
}