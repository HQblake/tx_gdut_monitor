package database

const (
	// MysqlDriver 数据库驱动
	MysqlDriver = "mysql"
	// RedisDriver 根据实际的情况选择驱动类型
	RedisDriver = "redis"
)

//Config 数据库配置结构体
type Config interface {
	GetDriver() string
	GetSource() string
}

