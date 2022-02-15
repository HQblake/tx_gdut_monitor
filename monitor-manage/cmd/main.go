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
	flag.StringVar(&configPath, "config", "./configs/config.yaml", "config path")
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
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
