package controller

import (
	R "../../TQStruct/base"
	TQ "../../TQBase/base"
	TQC "../../TQBase/constant"
	"encoding/json"
	"net/http"
	tNet "../../TQBase/network"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	url :=  TQC.LoginHost+ r.RequestURI
	http.Redirect(w,r,url,http.StatusMovedPermanently)
}

func Login(w http.ResponseWriter, r *http.Request)  {
	hc:=tNet.GetHttpClient(w,r)
	params:=hc.GetParam()
	logonVersion:=params.Get("lv")
	clientVersion:=params.Get("cv")
	//请求校验
	if logonVersion=="" {
		hc.ReturnMsg(R.ErrorMsg("请输入登录器版本号"))
		return
	}
	if clientVersion=="" {
		hc.ReturnMsg(R.ErrorMsg("请输入客户端版本号"))
		return
	}
	hc.ReturnMsg(R.OK())
}

func Register(w http.ResponseWriter, r *http.Request)  {
	TQ.LogErr(json.NewEncoder(w).Encode(R.Error().OutLog()))
}