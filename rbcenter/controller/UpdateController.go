package controller

import (
	R "../../rbstruct/base"
	"../../rbwork/constant"
	"net/http"
	"../../rbwork/network"
    "../service"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	//302重定向到新页面
	url := constant.MainHost + r.RequestURI
	http.Redirect(w,r,url,http.StatusMovedPermanently)
}

func CheckVersion(w http.ResponseWriter, r *http.Request)  {
	hc:=network.GetHttpClient(w,r)
	//获取最新版本信息
    versions :=service.CheckVersion()
	if len(versions)<=0 {
		hc.ReturnMsg(R.ErrorMsg("未查询到最新版本信息"))
	}
	hc.ReturnMsg(R.OK().SetData(versions))
}

