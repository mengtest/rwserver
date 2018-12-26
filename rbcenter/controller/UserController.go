package controller

import (
	"net/http"
	"../service"
	"../../rbwork/network"
	R "../../rbstruct/base"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	hc:=network.GetHttpClient(w,r)
	//获取最新版本信息
	versions :=service.CheckVersion()
	if len(versions)<=0 {
		hc.ReturnMsg(R.ErrorMsg("未查询到最新版本信息"))
	}
	hc.ReturnMsg(R.OK().SetData(versions))
}

