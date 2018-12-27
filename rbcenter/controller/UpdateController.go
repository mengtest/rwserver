package controller

import (
	R "../../rbstruct/base"
	"net/http"
	"../../rbwork/network"
    "../service"
)

func CheckVersion(w http.ResponseWriter, r *http.Request)  {
	hc:=network.GetHttpClient(w,r)
	//获取最新版本信息
    versions :=service.CheckVersion()
	if len(versions)<=0 {
		hc.ReturnMsg(R.ErrorMsg("未查询到最新版本信息"))
		return
	}
	for i,version :=range versions {
		versions[i]=service.GetNewVersion(version)
	}
	hc.ReturnMsg(R.OK().SetData(versions))
}


