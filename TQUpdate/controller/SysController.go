package controller

import (
	R "../../TQStruct/base"
	TQ "../../TQBase/base"
	TQC "../../TQBase/constant"
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	//302重定向到更新页面
	url :=  TQC.UPDATE_HOST+ r.RequestURI
	http.Redirect(w,r,url,http.StatusMovedPermanently)
}

func Update(w http.ResponseWriter, r *http.Request)  {
	logonVersion, err1 := r.URL.Query()["lv"]
	clientVersion, err2 := r.URL.Query()["cv"]
	//请求校验
	if !err1 || len(logonVersion) < 1 {
		TQ.LogErr(json.NewEncoder(w).Encode(R.ErrorMsg("请输入登录器版本号").OutLog()))
		return
	}
	if !err2 || len(clientVersion) < 1 {
		TQ.LogErr(json.NewEncoder(w).Encode(R.ErrorMsg("请输入客户端版本号").OutLog()))
		return
	}
	TQ.LogErr(json.NewEncoder(w).Encode(R.OK().OutLog()))

}

func FileMd5Check(w http.ResponseWriter, r *http.Request)  {
	fileName, err1 := r.URL.Query()["fileName"]
	md5, err2 := r.URL.Query()["md5"]
	//请求校验
	if !err1 || len(fileName) < 1 {
		TQ.LogErr(json.NewEncoder(w).Encode(R.ErrorMsg("请输入文件名").OutLog()))
		return
	}
	if !err2 || len(md5) < 1 {
		TQ.LogErr(json.NewEncoder(w).Encode(R.ErrorMsg("请输入md5值").OutLog()))
		return
	}
	TQ.LogErr(json.NewEncoder(w).Encode(R.OK().OutLog()))

}