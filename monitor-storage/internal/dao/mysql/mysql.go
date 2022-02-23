package mysql

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"time"
)

type Client struct {
	db       *sql.DB
	agents   *CientLock
	metricMx sync.Mutex
	metrics  sync.Map
}

func NewClient(s *MySQLSetting) *Client {
	//// 加载数据库驱动
	dns := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := sql.Open(s.DBType,
		fmt.Sprintf(dns, s.Username, s.Password, s.Host, s.DBName, s.Charset, s.ParseTime))

	if err != nil {
		log.Fatalln(err)
	}

	db.SetConnMaxLifetime(time.Duration(s.MaxLifetime) * time.Minute)
	db.SetMaxOpenConns(s.MaxOpenConns)
	db.SetMaxIdleConns(s.MaxIdleConns)
	defer log.Println("MySQL Connection Succeeded")

	// 初始化metric列表
	metrics := sync.Map{}
	row, _ := db.Query("SELECT * FROM metric")
	var id int
	var metric string
	for row.Next() {
		row.Scan(&id, &metric)
		metrics.Store(metric, id)
		log.Printf("Load metric: %d-%v", id, metric)
	}

	// 返回客户端
	return &Client{db, NewLock(), sync.Mutex{}, metrics}
}
