package http

// IShow 图表展示的方案设计，可由自青设计
type IShow interface {
	// 比如根据agentId和metric获取一天内的指标情况等
}

// Show 实现IShow接口的实例，包括数据处理，最后调用存储模块的rpc服务
type Show struct {

}