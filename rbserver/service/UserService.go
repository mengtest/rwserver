package service

import (
	"../../rbwork/network"
	"../util"
	"../../rbwork/base"
	R "../../rbstruct/base"
	)


//登录授权校验
func (s *Service) Login(tcpClient *network.TcpClient,umap map[string]interface{})  {
	strToken:=umap["token"].(string)
	requestId:=umap["requestId"].(string)
	claims,err :=base.DecodeToken(strToken)
	if err != nil {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("Login",requestId,"token无效")))
		return
	}
	if umap["mac"].(string) != claims["mac"].(string) {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("Login",requestId,"token无效,请先登录")))
		return
	}
	userId:=claims["uid"].(string)
	tcpClient.SetIsLogin(true)
	tcpClient.SetUserId(userId)
	tcpClient.SetMac(umap["mac"].(string))

	tcpClient.Write(base.Struct2Json(R.TcpOK("Login",requestId)))
}


//获取角色列表
func (s *Service) GetRoles(tcpClient *network.TcpClient,umap map[string]interface{})  {

}

//选择角色进入
func (s *Service) LoginRole(tcpClient *network.TcpClient,umap map[string]interface{})  {
	requestId:=umap["requestId"].(string)
	roleId:=umap["roleId"].(string)
	//load role info

	tcpClient.SetRoleId(roleId)

	util.Clients.Delete(tcpClient.GetIP()) //清除游客模式连接
	util.Clients.Set(tcpClient.GetUserId()+roleId,tcpClient)     //设置用户ID为主键
	tcpClient.Write(base.Struct2Json(R.TcpOK("LoginRole",requestId)))
}


