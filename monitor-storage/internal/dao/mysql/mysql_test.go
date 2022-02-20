package mysql

import (
	"fmt"
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
