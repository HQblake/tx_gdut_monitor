package mysql

import (
	"fmt"
	"gitee.com/zekeGitee_admin/tx_gdut_monitor/monitor-storage/internal/model"
	"log"
	"testing"
)

var MS = &MySQLSetting{
	"mysql",
	"root",
	"tx-monitor-gdut-group2",
	"81.71.165.211:3307",
	"monitor",
	"utf8mb4",
	true,
	10,
	10,
	3,
}

func TestClient_GetAllAgentInfo(t *testing.T) {
	client := NewClient(MS)
	agents := client.GetAllAgentInfo()
	for _, agent := range agents {
		fmt.Println(agent)
	}
}

func TestClient_GetAgentInfoByIPAndLocal(t *testing.T) {
	client := NewClient(MS)
	agent := client.GetAgentInfoByIPAndLocal("172.18.0.5", "beijing")
	fmt.Println(agent)
}

func TestClient_UpdateCheckConfig(t *testing.T) {
	client := NewClient(MS)
	id, err := client.UpdateCheckConfig(&model.CheckConfig{
		ID:        -1,
		IP:        "172.18.0.5",
		Local:     "beijing",
		Metric:    "mem_rate",
		Method:    2,
		Period:    "5m",
		Threshold: `{"0":"0.02","1":"0.03","2":"0.06","3":"0.08"}`,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
}

func TestClient_GetCheckConfigsByIPAndLocal(t *testing.T) {
	client := NewClient(MS)
	configs := client.GetCheckConfigsByIPAndLocal("172.18.0.5", "beijing")
	for _, config := range configs {
		fmt.Println(config)
	}
}
