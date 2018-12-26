package controller

import (
	"net/http"
	"../service"
	"../../rbwork/network"
	R "../../rbstruct/base"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	hc:=network.GetHttpClient(w,r)
	params:=hc.GetParam()

	mobile:=params.Get("mobile")
	checkCode:=params.Get("checkCode")
	if mobile=="" {
		hc.ReturnMsg(R.ErrorMsg("请输入手机号"))
		return
	}
	if checkCode=="" {
		hc.ReturnMsg(R.ErrorMsg("请输入短信验证码"))
		return
	}
	//获取最新版本信息
	versions :=service.CheckVersion()
	if len(versions)<=0 {
		hc.ReturnMsg(R.ErrorMsg("未查询到最新版本信息"))
	}
	hc.ReturnMsg(R.OK().SetData(versions))
}

func LoginPwd(w http.ResponseWriter, r *http.Request)  {
	hc:=network.GetHttpClient(w,r)
	params:=hc.GetParam()

	mobile:=params.Get("mobile")
	password:=params.Get("password")
	if mobile=="" {
		hc.ReturnMsg(R.ErrorMsg("请输入用户名"))
		return
	}
	if password=="" {
		hc.ReturnMsg(R.ErrorMsg("请输入密码"))
		return
	}
	//获取最新版本信息
	versions :=service.CheckVersion()
	if len(versions)<=0 {
		hc.ReturnMsg(R.ErrorMsg("未查询到最新版本信息"))
	}
	hc.ReturnMsg(R.OK().SetData(versions))
}

