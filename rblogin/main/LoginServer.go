package main

import (
	"net/http"
	"../route"
	TQ "../../rbwork/base"
)

func main() {
	TQ.Init(TQ.GetCurrentDirectory(),"LoginServer.log")
	router := route.NewRouter()
	TQ.LogInfo("login sever start...")
	TQ.LogFatal(http.ListenAndServe(":9091", router))
}
