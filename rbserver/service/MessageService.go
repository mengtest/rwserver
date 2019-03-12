package service

import (
	R "../../rbstruct/base"
	"../../rbstruct/net"
	"../../rbwork/base"
	"../../rbwork/constant"
	"../../rbwork/network"
	"../../rbwork/redis"
	Gloal "../util"
	"strconv"
)

//聊天消息
func (s *Service) Chat(tcpClient *network.TcpClient,msg string)  {
	req:=&net.ChatReq{}
	base.Json2Struct(msg,req)
	if !tcpClient.GetIsLogin() {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(req.Cmd,req.RequestId,"未授权")))
		return
	}


	switch req.NChannel {
	case 0:
		//世界
		c:=net.NewChat(tcpClient.GetRoleId(),tcpClient.GetRole().StrName,0,"",req.StrMsg,req.NChannel)
		for _,t:= range Gloal.Clients.GetMap() {
			t.Write(base.Struct2Json(R.TcpOK(req.Cmd,req.RequestId).SetData(c)))
		}
		return
	case 1:
		//当前
		go ChatToAroundPlayers(tcpClient,req.Cmd,req.RequestId,req.StrMsg,req.NChannel)
		return
	case 2:
		//地区
		return
	case 3:
		//组织
		return
	case 4:
		//队伍
		return
	case 5:
		//团队
		return
	case 6:
		//系统通知
		c:=net.NewChat(tcpClient.GetRoleId(),tcpClient.GetRole().StrName,0,"",req.StrMsg,req.NChannel)
		for _,t:= range Gloal.Clients.GetMap() {
			t.Write(base.Struct2Json(R.TcpOK(req.Cmd,req.RequestId).SetData(c)))
		}
		return
	case 10:
		//私聊
	    toRoleId:=strconv.FormatInt(req.LToRoleId,10)
		client:=Gloal.Clients.Get(toRoleId)
		c:=net.NewChat(tcpClient.GetRole().LId,tcpClient.GetRole().StrName,client.GetRole().LId,client.GetRole().StrName,msg,req.NChannel)
		client.Write(base.Struct2Json(R.TcpOK(req.Cmd,req.RequestId).SetData(c)))
		return
	default:
		return
	}
}

func ChatToAroundPlayers(tcpClient *network.TcpClient,cmd string,requestId string,msg string,channel int){
	role := tcpClient.GetRole()
	//---- 获取附近角色ID
	var roleIds []string
	for i := role.NChunkX - 1; i <= role.NChunkX+1; i++ {
		for j := role.NChunkY - 1; j <= role.NChunkY+1; j++ {
			rIds:= redis.Client.SMembers(constant.MapChunk +role.StrMapName+":"+ strconv.Itoa(i) + "#" + strconv.Itoa(j))
			roleIds = append(roleIds, rIds.Val() ...)
		}
	}
	//---- 发送信息
	for _, roleId := range roleIds {
		if roleId != "" {
			client:=Gloal.Clients.Get(roleId)
			c:=net.NewChat(role.LId,role.StrName,client.GetRole().LId,client.GetRole().StrName,msg,channel)
			client.Write(base.Struct2Json(R.TcpOK(cmd,requestId).SetData(c)))
		}
	}

}