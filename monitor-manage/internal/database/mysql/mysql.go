package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"tx_gdut_monitor/internal/database"
)

//InitConnection 初始化数据库连接
func InitConnection(config database.Config) (*sql.DB, error) {
	return getConnection(config)
}

func getConnection(config database.Config) (*sql.DB, error) {
	db, e := sql.Open(config.GetDriver(), config.GetSource())
	if e == nil {
		if err := db.Ping(); err != nil {
			return nil, err
		}
		db.SetMaxOpenConns(1000)
		db.SetMaxIdleConns(100)
		return db, nil
	}
	return nil, e
}

