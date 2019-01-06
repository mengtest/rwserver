package service

import (
	R "../../rbstruct/base"
	"../../rbstruct/net"
	"../../rbwork/base"
	"../../rbwork/constant"
	"../../rbwork/network"
	"../../rbwork/redis"
	"../util"
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

	switch req.ChatType {
	case 0:
		//世界发送消息
		for _,t:= range util.Clients.GetMap() {
			t.Write(base.Struct2Json(R.TcpOkMsg(req.Cmd,req.RequestId,req.Msg)))
		}
		return
	case 1:
		//当前频道向附近玩家发送消息
		go ChatToAroundPlayers(tcpClient,req.Cmd,req.RequestId,req.Msg)
		return
	case 10: //指定用户说话
		util.Clients.Get(strconv.FormatInt(req.ToRoleId,10)).Write(base.Struct2Json(R.TcpOkMsg(req.Cmd,req.RequestId,req.Msg)))
		return
	default:
		return
	}
}

func ChatToAroundPlayers(tcpClient *network.TcpClient,cmd string,requestId string,msg string){
	role := tcpClient.GetRole()
	var roleIds []string
	for i := role.NChunkX - 1; i <= role.NChunkX+1; i++ {
		for j := role.NChunkY - 1; j <= role.NChunkY+1; j++ {
			rIds:= redis.Client.SMembers(constant.MapChunk +role.StrMapName+":"+ strconv.Itoa(i) + "#" + strconv.Itoa(j))
			roleIds = append(roleIds, rIds.Val() ...)
		}
	}
	//---- 获取这些角色信息
	for _, roleId := range roleIds {
		if roleId != "" {
			client:=util.Clients.Get(roleId)
			client.Write(base.Struct2Json(R.TcpOkMsg(cmd,requestId,msg)))
		}
	}

}