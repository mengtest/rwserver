package controller

import (
	R "../../rbstruct/base"
	"../../rbwork/constant"
	"net/http"
	"../../rbwork/network"
    "../service"
)

func Index(w http.ResponseWriter, r *http.Request)  {
	//302重定向到更新页面
	url := constant.UpdateHost + r.RequestURI
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

func FileMd5Check(w http.ResponseWriter, r *http.Request)  {
	hc:=network.GetHttpClient(w,r)
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