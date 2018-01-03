package service

import (
	"reflect"
	"../../rbwork/network"
	"../util"
	"../../rbwork/base"
	R "../../rbstruct/base"
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
	case "1": //指定用户说话
		util.Clients.Get(umap["toUserId"].(string)).Write(base.Struct2Json(R.TcpOkMsg("chat",requestId,umap["msg"].(string))))
		return
	default:
		return
	}
}
