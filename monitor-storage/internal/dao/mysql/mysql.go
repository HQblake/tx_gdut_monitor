package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Client struct {
	db *sql.DB
}

func NewClient(s *MySQLSetting) *Client {
	// 加载数据库驱动
	dns := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := sql.Open(s.DBType,
		fmt.Sprintf(dns, s.UserName, s.PassWord, s.Host, s.DBName, s.Charset, s.ParseTime))

	if err != nil {
		log.Fatalln(err)
	}

	db.SetConnMaxLifetime(time.Duration(s.MaxLifetime) * time.Minute)
	db.SetMaxOpenConns(s.MaxOpenConns)
	db.SetMaxIdleConns(s.MaxIdleConns)
	return &Client{db}
}