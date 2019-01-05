package service

import (
	R "../../rbstruct/base"
	"../../rbwork/base"
	"../../rbwork/constant"
	"../../rbwork/network"
	"../../rbwork/redis"
	"../dao"
	"../util"
	"reflect"
	"strconv"
)

//登录授权校验
func (s *Service) Login(tcpClient *network.TcpClient, umap map[string]interface{}) {
	strToken := umap["token"].(string)
	requestId := umap["requestId"].(string)
	claims, err := base.DecodeToken(strToken)
	if err != nil {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("Login", requestId, "token无效")))
		return
	}
	mac := umap["mac"]
	if mac == nil || reflect.TypeOf(mac).String() != "string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("无效MAC")))
		return
	}

	if claims["mac"] == nil || umap["mac"].(string) != claims["mac"].(string) {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("Login", requestId, "token无效,请先登录")))
		return
	}
	userId := claims["uid"].(string)
	tcpClient.SetIsLogin(true)
	tcpClient.SetUserId(userId)
	tcpClient.SetMac(umap["mac"].(string))

	tcpClient.Write(base.Struct2Json(R.TcpOK("Login", requestId)))
}

//获取角色列表
func (s *Service) GetRoles(tcpClient *network.TcpClient, umap map[string]interface{}) {
	requestId := umap["requestId"].(string)
	if !tcpClient.GetIsLogin() {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("Login", requestId, "未授权，请先登录")))
		return
	}
	roles := dao.GetRolesByUserId(tcpClient.GetUserId())
	tcpClient.Write(base.Struct2Json(R.TcpOK("GetRoles", requestId).SetData(roles).OutLog()))
}

//选择角色进入
func (s *Service) LoginRole(tcpClient *network.TcpClient, umap map[string]interface{}) {
	requestId := umap["requestId"].(string)
	roleId := umap["roleId"]
	if !tcpClient.GetIsLogin() {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("Login", requestId, "未授权，请先登录")))
		return
	}
	if roleId == nil || reflect.TypeOf(roleId).String() != "string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("roleId参数无效")))
		return
	}
	//load role info
	role := dao.GetRoleByRoleId(roleId.(string))
	if role.LId <= 0 {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("角色不存在")))
		return
	}

	tcpClient.SetRoleId(roleId.(string))
	tcpClient.SetRole(role)
	redis.Sadd(constant.MAP_CHUNK+role.StrMapName+":"+strconv.Itoa(role.NChunkX)+":"+strconv.Itoa(role.NChunkY), roleId.(string))
	util.Clients.Delete(tcpClient.RemoteAddr())  //清除游客模式连接
	util.Clients.Set(roleId.(string), tcpClient) //设置角色ID为主键
	tcpClient.Write(base.Struct2Json(R.TcpOK("LoginRole", requestId).SetData(role).OutLog()))
}
