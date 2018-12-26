package main

import (
	"net/http"
	"../route"
	TQ "../../rbwork/base"
	"../../rbwork/db"
)


// 初始化参数
func init() {
	//初始化日志输出
	TQ.Init(TQ.GetCurrentDirectory(),"CenterServer.log")
	//初始化DB
	db.Init("root:123456@tcp(127.0.0.1:3306)/tianqi?charset=utf8")
}

func main() {
	//启动服务
	router := route.NewRouter()
	TQ.LogInfo("updateSever start...")
	TQ.LogErr(http.ListenAndServe(":9092", router))
}
