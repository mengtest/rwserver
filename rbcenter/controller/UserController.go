package controller

import (
	"net/http"
	"../service"
	"../../rbwork/network"
	R "../../rbstruct/base"
	"../../rbwork/base"
)


func Login(w http.ResponseWriter, r *http.Request)  {
	hc:=network.GetHttpClient(w,r)
	params:=hc.GetParam()

	strName:=params.Get("strName")
	strPwd:=params.Get("strPwd")

	if strName=="" {
		hc.ReturnMsg(R.ErrorMsg("请输入用户名"))
		return
	}
	if strPwd=="" {
		hc.ReturnMsg(R.ErrorMsg("请输入密码"))
		return
	}
	user,count:=service.GetUserByName(strName)
	if count<=0 {
		hc.ReturnMsg(R.ErrorMsg("用户不存在"))
		return
	}
	token:=base.CreateToken(string(user.LId))
	//获取最新版本信息
	user,ret,msg :=service.Login(strName,strPwd)
	if ret<0 {
		hc.ReturnMsg(R.ErrorMsg(msg))
	}
	hc.ReturnMsg(R.OK().SetData(token))
}

