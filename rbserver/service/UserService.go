package service

import (
	R "../../rbstruct/base"
	"../../rbstruct/net"
	"../../rbwork/base"
	"../../rbwork/constant"
	"../../rbwork/network"
	"../../rbwork/redis"
	"../../rbstruct/user"
	"../dao"
	Gloal "../util"
	"strconv"
)

//登录授权校验
func (s *Service) Login(tcpClient *network.TcpClient, msg string) {
	req:=&net.LoginReq{}
	base.Json2Struct(msg,req)

	claims, err := base.DecodeToken(req.Token)
	if err != nil {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(req.Cmd, req.RequestId, "token无效")))
		return
	}
	if claims["mac"] == nil || req.Mac != claims["mac"].(string) {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(req.Cmd,  req.RequestId, "token无效,请先登录")))
		return
	}
	userId := claims["uid"].(string)
	lUserId, err := strconv.ParseInt(userId, 10, 64)
	tcpClient.SetIsLogin(true)
	tcpClient.SetUserId(lUserId)
	tcpClient.SetMac(req.Mac)

	tcpClient.Write(base.Struct2Json(R.TcpOK(req.Cmd, req.RequestId)))
}

//获取角色列表
func (s *Service) GetRoles(tcpClient *network.TcpClient, msg string) {
	req:=&net.Req{}
	base.Json2Struct(msg,req)
	if !tcpClient.GetIsLogin() {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(req.Cmd, req.RequestId, "未授权，请先登录")))
		return
	}
	roles := dao.GetRolesByUserId(tcpClient.GetUserId())
	tcpClient.Write(base.Struct2Json(R.TcpOK(req.Cmd, req.RequestId).SetData(roles).OutLog()))
}

//选择角色进入
func (s *Service) LoginRole(tcpClient *network.TcpClient, msg string) {
	req:=&net.LoginRoleReq{}
	base.Json2Struct(msg,req)
	if !tcpClient.GetIsLogin() {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(req.Cmd, req.RequestId, "未授权，请先登录")))
		return
	}
	if req.RoleId <=0 {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("roleId参数无效")))
		return
	}
	//加载角色信息
	role := dao.GetRoleByRoleId(req.RoleId)
	if role.LId <= 0 {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("角色不存在")))
		return
	}
	//加载角色技能
	skills := dao.GetRoleSkillByRoleId(req.RoleId)
	role.Skills=skills
	role.Buffs=[]user.RoleBuff{}
	role.Action="idle0"
    //设置client角色信息
	tcpClient.SetRoleId(req.RoleId)
	tcpClient.SetRole(role)
	//设置角色所在地图块
	redis.Client.SAdd(constant.MapChunk+role.StrMapName+":"+strconv.Itoa(role.NChunkX)+"#"+strconv.Itoa(role.NChunkY), req.RoleId)
    //全局存储client
    strRoleId:=strconv.FormatInt(req.RoleId,10)
	Gloal.Clients.Delete(tcpClient.RemoteAddr())  //清除游客模式连接
	Gloal.Clients.Set(strRoleId, tcpClient) //设置角色ID为主键
    //返回自己的角色信息
	tcpClient.Write(base.Struct2Json(R.TcpOK(req.Cmd, req.RequestId).SetData(role).OutLog()))
    //向周围玩家同步下发自己的信息
	SyncPlayerToAroundPlayers(strRoleId,role,nil)
    //告知好友上线了

}

//创建角色
func (s *Service) CreateRole(tcpClient *network.TcpClient, msg string)  {
	req:=&net.CreateRoleReq{}
	base.Json2Struct(msg,req)
	if !tcpClient.GetIsLogin() {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(req.Cmd, req.RequestId, "未授权，请先登录")))
		return
	}
	roleInfo:=&user.RoleInfo{}
	roleInfo.LUserId=tcpClient.GetUserId()
	roleInfo.StrName=req.StrName
	roleInfo.NSex=req.NSex
	roleInfo.NOccId=req.NOccId
	roleInfo.StrOccName=req.StrOccName
	//新建角色默认出生在新手村
	roleInfo.StrMapName=constant.MAP_NAME
	roleInfo.FPosX=constant.ROLE_POSX
	roleInfo.FPosY=constant.ROLE_POSY
	roleInfo.FPosZ=constant.ROLE_POSZ
	roleInfo.FDirX=constant.ROLE_DIRX
	roleInfo.FDirY=constant.ROLE_DIRY
	roleInfo.FDirZ=constant.ROLE_DIRZ
	roleInfo.NChunkX=constant.ROLE_CHUNKX
	roleInfo.NChunkY=constant.ROLE_CHUNKY
	success:=dao.CreateRole(roleInfo)
	if success {
		tcpClient.Write(base.Struct2Json(R.TcpOK(req.Cmd, req.RequestId).SetData(roleInfo).OutLog()))
	}else{
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(req.Cmd, req.RequestId, "创建角色失败，请重试")))
	}
}