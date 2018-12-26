package main

import (
	"net/http"
	TQ "../../rbwork/base"
	"os"
	"../rbutil"
	"../route"
)



func init()  {
	TQ.Init(TQ.GetCurrentDirectory(),"FileServer.log")
	exist, _ := PathExists(rbutil.FilePath)
	if !exist {
		os.Mkdir(rbutil.FilePath,os.ModePerm)
	}
	rbutil.StaticHandler = http.StripPrefix("/download", http.FileServer(http.Dir(rbutil.FilePath)))
	TQ.LogInfo(rbutil.FilePath)
}

func main() {
	router := route.NewRouter()
	TQ.LogInfo("file sever start...")
	TQ.LogFatal(http.ListenAndServe(":9091", router))
}


// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}