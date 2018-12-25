package controller

import (
	R "../../TQStruct/base"
	TQC "../../TQBase/constant"
	"net/http"
	tNet "../../TQBase/network"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	//302重定向到更新页面
	url :=  TQC.UpdateHost+ r.RequestURI
	http.Redirect(w,r,url,http.StatusMovedPermanently)
}

func Update(w http.ResponseWriter, r *http.Request)  {
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

func FileMd5Check(w http.ResponseWriter, r *http.Request)  {
	hc:=tNet.GetHttpClient(w,r)
	params:=hc.GetParam()
	fileName:=params.Get("fileName")
	md5:=params.Get("md5")
	//请求校验
	if fileName=="" {
		hc.ReturnMsg(R.ErrorMsg("请输入文件名"))
		return
	}
	if md5 == "" {
		hc.ReturnMsg(R.ErrorMsg("请输入md5值"))
		return
	}
	hc.ReturnMsg(R.OK())
}