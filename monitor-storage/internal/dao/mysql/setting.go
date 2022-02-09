package mysql

type MySQLSetting struct {
	RunMode         string
	UserName        string
	PassWord        string
	Host            string
	DBName          string
	Charset         string
	ParseTime       bool
	MaxIdleConns    int
	MaxOpenConns    int
	SingularTable   bool
	SkipTransaction bool
}
