### 发送模块需求梳理

##### 开发者：吴泽楷

##### 日期： 202201

#### 对接判定服务

1、提供interface类型供判定服务调用
2、interface类型提供用于对判定服务发来的告警信息进行处理的方法
3、interface类型暂定如下：
```go
type ISend interface {
	Send(alert model.AlertInfo) error
}
```

#### 对接管理服务/服务监控系统
1、基于GRPC与管理服务进行交互
2、主要提供的grpc服务包括
- 更新/新增/删除配置处理，监听管理服务发来配置更新，及时更新缓存配置
- 初始化配置处理，用于项目初始化时，管理服务提供项目的初始化发送配置
- 配置检查处理，避免脏数据入库

#### 支持告警收敛算法

实际监控过程中，可能出现海量告警的情况，此时运维人员可能因为过多的报警信息而感觉“麻木”，从而忽略了重要的告警信息，也就是“狼来了”的类比情况。基于此，系统告警过程中支持使用告警收敛处理，用户可根据自己的实际需求进行调整，支持的告警类型如下：

- 即时告警，发送模块收到告警信息即刻发送告警信息，适合告警阈值较高，告警频率较高的场景
- 聚合告警，根据告警周期的设置情况，以一定的时间频率发送周期时间内的触发的告警消息，统一整合成一封邮件/消息发出，避免海量告警，适合告警频率较高的场景
- 滚动收敛告警，根据告警周期的设置情况，以一定的时间频率发送周期时间内的触发的告警消息，发送的告警信息会基于滚动收敛算法进行过滤，该算法待完善，同样适合告警频率较高的场景

**配置详情：**

```yaml
Alert:
  Convergence: 1 # 告警收敛方式，0为即时告警，1为聚合收敛，2为滚动收敛(未完善)，默认为0
  Interval: 1 # 告警周期，单位为min，收敛方式为1或2时有效，默认为1min
```

#### 支持告警等级
- info-提示 0
- warn-警告 1
- Error-严重 2
- panic-致命 3

#### 支持告警对象
- 邮件告警
- nsq mq
- kafka mq
- http接口告警

#### 核心代码接口定义

##### 输出器

上述每一个告警对象都是一个IOutput输出器，开发过程中可以根据实际需求动态扩增所需的输出器细节

```golang
// IOutput 输出器
type IOutput interface {
    // 告警等级
	Level() Level
	// Reset 更新配置
	Reset(level Level, config interface{}) error
	// Output 内容输出
	Output(infos []model.Info) error
	// Finish 结束，即删除配置后的析构处理
	Finish() error
}
```

**输出管理器**

在输出管理器中维护多个IOutput输出器，根据实际的业务需要对多个IOutput进行增删改，同时控制IOutput的Reset，Output等操作

```golang
type IOutputs interface {
	ID() string
	Get(id int) (IOutput, bool)
	Del(id int) error
	Set(id int, conf Config) error
	List() []IOutput
	Output(infos []model.Info) error
	Check(conf Config) error
}
```

