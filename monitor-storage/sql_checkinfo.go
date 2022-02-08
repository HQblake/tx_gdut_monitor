package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

var dbCheckinfo *sql.DB //连接池对象

func initDbCheckinfo() (err error) {
	// DSN:Data Source Name
	//dsn := "root:123456@tcp(127.0.0.1:3306)/test2"
	dsn := "root:dajiao38ma@/jiaotest"
	dbCheckinfo, err = sql.Open("mysql", dsn) //open不会校验用户名和密码是否正确
	if err != nil {
		return
	}
	err = dbCheckinfo.Ping()
	if err != nil {
		return
	}

	dbCheckinfo.SetMaxOpenConns(10) //设置数据库连接池的最大连接数 10
	dbCheckinfo.SetMaxIdleConns(5)  //设置最大空闲连接数
	return
}

type checkinfo struct { //结构体
	id   int
	ip	string
	port string
	local string
	metric string
	method int
	period string
	threshold string

}
func insertCheckinfo(u2 checkinfo) {
	rows, err := dbCheckinfo.Query("INSERT INTO checkinfo (ip,port,local,metric,method,period,threshold)\n" +
		"VALUES(?,?,?,?,?,?,?)",u2.ip,u2.port,u2.local,u2.metric,u2.method,u2.period,u2.threshold)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
}
func updateCheckinfo(u3 checkinfo) {
	rows, err := dbCheckinfo.Query("update checkinfo set metric=?,method=?,period=?,threshold=? where id=?",
		u3.metric,u3.method,u3.period,u3.threshold,u3.id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("更新")
	defer rows.Close()
}
func selectCheckinfo(id int) {
	var u1 checkinfo
	//// 1.写查询单条记录的sql语句
	//sqlStr := "select id,ip,port,local,metric,method,period,threshold from checkinfo where id <?;"
	////2.执行并拿到结果
	////必须对row对象调用scan方法，该方法会释放数据库连接
	//db_checkinfo.QueryRow(sqlStr, id).Scan(&u1.id, &u1.ip, &u1.port, &u1.local, &u1.metric, &u1.method, &u1.period, &u1.threshold) //从连接池拿一个连接出去数据库查询单条记录
	////打印结果
	//fmt.Println(u1)

	//查询d多条数据，指定字段名，返回sql.Rows结果集
	rows, err := dbCheckinfo.Query("select id,ip,port,local,metric,method,period,threshold from checkinfo where id <?;",id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&u1.id, &u1.ip, &u1.port, &u1.local, &u1.metric, &u1.method, &u1.period, &u1.threshold)
		fmt.Println(u1)
	}
	defer rows.Close()
}

func main() {
	err := initDbCheckinfo()
	if err != nil {
		fmt.Printf("init DB failed err:%v\n", err)
	}
	fmt.Println("连接数据库成功")

	selectCheckinfo(30)
	var u2 checkinfo
	u2 =checkinfo{
		0,
		"uiyi",
		"7890",
		"hkjh",
		"jkhk",
		3,
		"kjk",
		"hkjh",
	}
	insertCheckinfo(u2)
	selectCheckinfo(30)
	var u3 checkinfo
	u3 =checkinfo{
		2,
		"uiyi",
		"7890",
		"hkjh",
		"jkhk",
		3,
		"kjk",
		"hkjh",
	}
	updateCheckinfo(u3)
	selectCheckinfo(30)
}

