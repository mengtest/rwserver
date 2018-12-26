package main

import (
	"net/http"
	"../route"
	TQ "../../rbwork/base"
	"../../rbwork/db"
)

func main() {
	//初始化日志输出
	TQ.Init(TQ.GetCurrentDirectory(),"UpdateServer.log")
	//初始化路由
	router := route.NewRouter()
	//初始化DB
	db.Init("root:123456@tcp(127.0.0.1:3306)/tianqi?charset=utf8")
	//启动服务
	TQ.LogInfo("updateSever start...")
	TQ.LogErr(http.ListenAndServe(":9092", router))
}
