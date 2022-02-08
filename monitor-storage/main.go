package main

import (
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/global/service"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {

}

func main() {
	server := grpc.NewServer()
	metricService, err := service.NewMetricService()
	if err != nil {
		log.Fatalln(err)
	}
	manageService, err := service.NewManageService()
	if err != nil {
		log.Fatalln(err)
	}
	proto.RegisterMetricServiceServer(server, metricService)
	proto.RegisterManageConfigServer(server, manageService)
	lis, _ := net.Listen("tcp", ":")
	server.Serve(lis)
}
