package main

import (
	"flag"
	"log"
	"net"

	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/configs"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-manage/internal/service/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var configPath string
var debug bool

func init() {
	flag.StringVar(&configPath, "config", "../configs/config.yaml", "config path")
	flag.BoolVar(&debug, "debug", true, "is debug")
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
	l, err := net.Listen("tcp", configs.GetAdminServerAddr())
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	// 初始化管理服务
	admin := service.Register(alertConn, storeConn)

	// 注册grpc服务
	admin.RegisterService(l)

	// 开启监听
	log.Fatal(router(admin.ApiHandler).Run(configs.GetAdminListenAddr()))
}

func router(h *http.Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(Cors())
	r.Use(gin.Logger())
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	// agent
	agent := r.Group("/agent")
	agent.GET("/list", h.GetAllAgent)
	agent.GET("/sendList", h.GetAllAgentSendInfo)
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
	show.GET("/metrics/:ip/:local/:metric/:begin/:limit", h.GetMetricsWithTime)
	show.GET("/warnList", h.GetWarnInfo)
	show.GET("/warnId/:id", h.GetWarnInfoWithId)
	show.GET("/warnParams/:ip/:local/:metric/:level/:start/:end", h.GetWarnInfoWithParams)
	show.POST("/del/:id", h.DelWarnInfo)
	return r
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 正式环境最好限制源
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.JSON(200, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
