package db

/**
 * Title:mysql连接池
 * User: iuoon
 * Date: 2018-12-24
 * Version: 1.0
 */

import (
"database/sql"
TQ "../base"
)

var MySQL *sql.DB

func InitMySQL(dataSourceName string) {
	//dataSourceName=root:@tcp(127.0.0.1:3306)/tianqi?charset=utf8
	var err error
	MySQL, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		TQ.LogError("数据初始化连接失败",err)
		return
	}
	MySQL.SetMaxOpenConns(1000)
	MySQL.SetMaxIdleConns(10)
	MySQL.Ping()
}
