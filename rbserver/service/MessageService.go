package service

import (
	"reflect"
	"../../rbwork/network"
	"../util"
	"../../rbwork/base"
	R "../../rbstruct/base"
	"strconv"
	"../../rbwork/constant"
	"../../rbwork/redis"
)

//聊天消息
func (s *Service) Chat(tcpClient *network.TcpClient,umap map[string]interface{})  {
	requestId:=umap["requestId"].(string)
	if !tcpClient.GetIsLogin() {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("chat",requestId,"未授权")))
		return
	}
	ty0:=reflect.TypeOf(umap["chatType"]).String()
	ty1:=reflect.TypeOf(umap["toUserId"]).String()
	ty2:=reflect.TypeOf(umap["msg"]).String()
	if ty0!="string" || ty1 !="string" || ty2 !="string"{
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("chat",requestId,"请求参数错误")))
		return
	}

	switch umap["chatType"].(string) {
	case "0":
		//世界发送消息
		for _,t:= range util.Clients.GetMap() {
			t.Write(base.Struct2Json(R.TcpOkMsg("chat",requestId,umap["msg"].(string))))
		}
		return
	case "1":
		//当前频道向附近玩家发送消息
		go ChatToAroundPlayers(tcpClient,requestId,umap["msg"].(string))
		return
	case "10": //指定用户说话
		util.Clients.Get(umap["toUserId"].(string)).Write(base.Struct2Json(R.TcpOkMsg("chat",requestId,umap["msg"].(string))))
		return
	default:
		return
	}
}

func ChatToAroundPlayers(tcpClient *network.TcpClient,requestId string,msg string){
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
			client.Write(base.Struct2Json(R.TcpOkMsg("chat",requestId,msg)))
		}
	}

}