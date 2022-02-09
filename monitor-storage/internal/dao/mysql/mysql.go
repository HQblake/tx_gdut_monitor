package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

type Client struct {
	db *gorm.DB
}

func NewClient(s *MySQLSetting) *Client {
	// 打印SQL日志
	var lo logger.Interface
	if s.RunMode == "debug" {
		lo = logger.Default.LogMode(logger.Silent)
	} else {
		lo = logger.Discard.LogMode(logger.Silent)
	}

	// 加载数据库驱动
	dns := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(mysql.Open(
		fmt.Sprintf(dns, s.UserName, s.PassWord, s.Host, s.DBName, s.Charset, s.ParseTime)),
		&gorm.Config{
			SkipDefaultTransaction: s.SkipTransaction,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: s.SingularTable,
			},
			Logger: lo,
		})

	if err != nil {
		log.Fatalln(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(s.MaxIdleConns)
	sqlDB.SetMaxOpenConns(s.MaxOpenConns)
	return &Client{db}
}
