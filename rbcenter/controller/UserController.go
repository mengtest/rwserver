package controller

import (
	R "../../rbstruct/base"
	"../../rbwork/base"
	"../../rbwork/network"
	"../service"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	hc := network.GetHttpClient(w, r)
	params := hc.GetParam()

	strName := params.Get("strName")
	strPwd := params.Get("strPwd")
	strMac := params.Get("strMac")

	if strName == "" {
		hc.ReturnMsg(R.ErrorMsg("请输入用户名"))
		return
	}
	if strPwd == "" {
		hc.ReturnMsg(R.ErrorMsg("请输入密码"))
		return
	}
	user, count := service.GetUserByName(strName)
	if count <= 0 {
		hc.ReturnMsg(R.ErrorMsg("用户不存在"))
		return
	}
	//校验密码
	pwd := base.GetMd5(base.DesEncode(strPwd))
	if pwd != user.StrPwd {
		hc.ReturnMsg(R.ErrorMsg("密码错误"))
		return
	}
	//生成token
	token := base.CreateToken(string(user.LId), strMac)

	userData := service.UserData{}
	userData.User.LId=user.LId
	userData.User.StrName=user.StrName
	userData.Token = token
	hc.ReturnMsg(R.OK().SetData(userData))
}

func Register(w http.ResponseWriter, r *http.Request) {
	hc := network.GetHttpClient(w, r)
	params := hc.GetParam()

	strName := params.Get("strName")
	strPwd := params.Get("strPwd")
	strMac := params.Get("strMac")

	if strName == "" {
		hc.ReturnMsg(R.ErrorMsg("请输入用户名"))
		return
	}
	if strPwd == "" {
		hc.ReturnMsg(R.ErrorMsg("请输入密码"))
		return
	}
	_, count := service.GetUserByName(strName)
	if count > 0 {
		hc.ReturnMsg(R.ErrorMsg("用户已存在"))
		return
	}
	user := service.User{}
	user.StrName = strName
	user.StrPwd = base.GetMd5(base.DesEncode(strPwd))
	ret := service.SaveUser(user)
	if ret <= 0 {
		hc.ReturnMsg(R.ErrorMsg("注册失败，请重新注册"))
		return
	}
	//校验密码

	//生成token
	token := base.CreateToken(strconv.FormatInt(user.LId, 10), strMac)

	userData := service.UserData{}
	userData.User.LId=user.LId
	userData.User.StrName=user.StrName
	userData.Token = token
	hc.ReturnMsg(R.OK().SetData(userData))
}
