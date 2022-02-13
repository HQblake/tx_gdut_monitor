package metrics

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func GetTimeStamp() int64 { //获取本机当前的时间戳
	unixtime := time.Now().Unix()   //获取当前的时间戳 （秒）
	fmt.Println("当前时间戳：", unixtime) //当前时间戳： 1587894706
	return unixtime
}

func GetCpuRate() float64 { //获取当前本机的Cpu使用率
	//res, err := cpu.Times(false)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//return ((res[0].Total() - res[0].Idle) / res[0].Total()) * 100
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemRate() float64 { //获取当前本机的内存使用率
	v, _ := mem.VirtualMemory()
	return v.UsedPercent
}
