package handle

import (
	R "../../rbstruct/base"
	"../../rbwork/base"
	"../../rbwork/network"
	"reflect"
	"../util"
	"time"
)

func HandleMsg(tcpClient *network.TcpClient,msg string)  {
	umap,err:=base.Json2map(msg)
	if err !=nil {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("无效请求")))
		return
	}
	cmd:=umap["cmd"]
	requestId:=umap["requestId"]
	type1:=reflect.TypeOf(cmd).String()
	if cmd ==nil || type1 !="string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("无效CMD")))
		return
	}
	type2:=reflect.TypeOf(requestId).String()
	if requestId ==nil ||  type2 !="string"{
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("无效ID")))
		return
	}

	switch cmd {
	case Login:
		login(tcpClient,umap)
		return
	case Chat:
		chat(tcpClient,umap)
		return
	case Move:
		move(tcpClient,umap)
		return
	case Ping:
		ping(tcpClient)
		return
	default:
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(cmd.(string),requestId.(string),"无效请求")))
		return
	}

}

//登录授权校验
func login(tcpClient *network.TcpClient,umap map[string]interface{})  {
	strToken:=umap["token"].(string)
	requestId:=umap["requestId"].(string)
	claims,err :=base.DecodeToken(strToken)
	if err != nil {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("login",requestId,"token无效")))
		return
	}
	if umap["mac"].(string) != claims["mac"].(string) {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("login",requestId,"token无效,请先登录")))
		return
	}
	tcpClient.SetIsLogin(true)

	userId:=claims["uid"].(string)

	util.Clients.Delete(tcpClient.GetIP()) //清除游客模式连接
	util.Clients.Set(userId,tcpClient)     //设置用户ID为主键

	tcpClient.Write(base.Struct2Json(R.TcpOK("login",requestId)))
}

//聊天消息
func chat(tcpClient *network.TcpClient,umap map[string]interface{})  {
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

//移动
func move(tcpClient *network.TcpClient,umap map[string]interface{})  {

}

//攻击敌人

//接任务

//完成任务

//升级

//获取物品（装备、物品）


func ping(tcpClient *network.TcpClient)  {
	tcpClient.SetTime(time.Now().Unix())
}