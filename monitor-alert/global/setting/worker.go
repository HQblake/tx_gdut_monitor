package setting

type WorkersSetting struct {
	Capacity         int    // 协程池容量
	ExpiryDuration   string // worker 协程过期时间
	Nonblocking      bool   // 是否启用非阻塞提交
	PreAlloc         bool   // 是否预分配内存
	MaxBlockingTasks int    // 最大阻塞任务数
}
