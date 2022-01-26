package admin

// 接口化定义，定义抽象的数据库操作方法

// IOutput 发送模块配置的数据库操作方法
type IOutput interface {
	GetConfig() error
	AddNewConfig() error
	UpdateConfig() error
	DelConfig() error
}
// ICheck 判定模块配置的数据库操作方法
type ICheck interface {

}
