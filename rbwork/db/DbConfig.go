package db

/**
 * Title:mysql连接池
 * User: iuoon
 * Date: 2018-12-24
 * Version: 1.0
 */

import (
     "../base"
	 "github.com/jmoiron/sqlx"
	_"github.com/go-sql-driver/mysql"  //注意驱动包要引入
)

var DB *sqlx.DB

func Init(dataSourceName string) {
	//dataSourceName=root:@tcp(127.0.0.1:3306)/tianqi?charset=utf8
	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		base.LogError("数据初始化连接失败",err)
		return
	}
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(10)
	DB=db
	base.LogInfo("connect database success")
}
