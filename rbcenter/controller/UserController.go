package controller

import (
	"net/http"
	"../service"
	"../../rbwork/network"
	R "../../rbstruct/base"
	"../../rbwork/redis"
	"../../rbwork/constant"
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
	rck := redis.GetStringValue(constant.SMS_MOBILE+mobile)
	if  rck=="" || rck !=checkCode {
		hc.ReturnMsg(R.ErrorMsg("验证码无效"))
		return
	}
	//获取最新版本信息
	user,err :=service.GetUserByMobile(mobile)
	if err !=nil {
		hc.ReturnMsg(R.ErrorMsg("用户不存在"))
	}
	hc.ReturnMsg(R.OK().SetData(user))
}

func LoginPwd(w http.ResponseWriter, r *http.Request)  {
	hc:=network.GetHttpClient(w,r)
	params:=hc.GetParam()

	mobile:=params.Get("mobile")
	pwd:=params.Get("pwd")
	if mobile=="" {
		hc.ReturnMsg(R.ErrorMsg("请输入用户名"))
		return
	}
	if pwd=="" {
		hc.ReturnMsg(R.ErrorMsg("请输入密码"))
		return
	}
	//获取最新版本信息
	user,ret,msg :=service.Login(mobile,pwd)
	if ret<0 {
		hc.ReturnMsg(R.ErrorMsg(msg))
	}
	hc.ReturnMsg(R.OK().SetData(user))
}

