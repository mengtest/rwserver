package main

import (
	"net/http"
	"../route"
	TQ "../../TQBase/base"
)

func main() {
	TQ.Init(TQ.GetCurrentDirectory(),"UpdateServer.log")
	router := route.NewRouter()
	TQ.LogInfo("updateSever start...")
	TQ.LogErr(http.ListenAndServe(":9092", router))
}
