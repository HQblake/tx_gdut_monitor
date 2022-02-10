package mysql

type MySQLSetting struct {
	DBType          string
	UserName        string
	PassWord        string
	Host            string
	DBName          string
	Charset         string
	ParseTime       bool
	MaxIdleConns    int
	MaxOpenConns    int
	MaxLifetime     int
	SingularTable   bool
	SkipTransaction bool
}