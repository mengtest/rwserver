package main

import (
	"net/http"
	"../route"
	TQ "../../TQBase/base"
)

func main() {
	TQ.Init(TQ.GetCurrentDirectory(),"LoginServer.log")
	router := route.NewRouter()
	TQ.LogInfo("login sever start...")
	TQ.LogFatal(http.ListenAndServe(":8080", router))
}
