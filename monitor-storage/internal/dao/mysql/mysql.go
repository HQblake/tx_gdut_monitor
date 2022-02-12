package mysql

import (
	"database/sql"
	"fmt"

	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"time"
)

type Client struct {
	db *sql.DB
}

func NewClient(s *MySQLSetting) *Client {
	//// 加载数据库驱动
	//dns := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	//db, err := sql.Open(s.DBType,
	//	fmt.Sprintf(dns, s.UserName, s.PassWord, s.Host, s.DBName, s.Charset, s.ParseTime))
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//db.SetConnMaxLifetime(time.Duration(s.MaxLifetime) * time.Minute)
	//db.SetMaxOpenConns(s.MaxOpenConns)
	//db.SetMaxIdleConns(s.MaxIdleConns)

	// DSN:Data Source Name
	//dsn := "root:123456@tcp(127.0.0.1:3306)/test2"
	var dbCheckinfo Client
	var err error

	dsn :=s.Username+":"+s.Password+"@/"+s.DBName
	dbCheckinfo.db, err = sql.Open("mysql", dsn) //open不会校验用户名和密码是否正确
	if err != nil {
		log.Fatalln(err)
	}
	err =dbCheckinfo.db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	dbCheckinfo.db.SetMaxOpenConns(s.MaxOpenConns) //设置数据库连接池的最大连接数 10
	dbCheckinfo.db.SetMaxIdleConns(s.MaxOpenConns)  //设置最大空闲连接数
	fmt.Println("连接数据库成功")
	return &Client{dbCheckinfo.db}
}
