package controller

import (
	R "../../rbstruct/base"
	"net/http"
	"../../rbwork/network"
    "../service"
	"fmt"
)

func CheckVersion(w http.ResponseWriter, r *http.Request)  {
	hc:=network.GetHttpClient(w,r)
	//获取最新版本信息
    versions :=service.CheckVersion()
    user,_ := service.GetUserByMobile("1111")
    fmt.Println(user)
	if len(versions)<=0 {
		hc.ReturnMsg(R.ErrorMsg("未查询到最新版本信息"))
		return
	}
	hc.ReturnMsg(R.OK().SetData(versions))
}

