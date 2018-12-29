package handle

import (
	R "../../rbstruct/base"
	"../../rbwork/base"
	"../../rbwork/network"
	"github.com/goinggo/mapstructure"
	"reflect"
	"../util"
)

type LoginStruct struct {
   Cmd string
   Token string
   Mac string
}

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
	default:
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(requestId.(string),"无效请求")))
		return
	}

}

func login(tcpClient *network.TcpClient,umap map[string]interface{})  {
	strToken:=umap["token"].(string)
	requestId:=umap["requestId"].(string)
	claims,err :=base.DecodeToken(strToken)
	if err != nil {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(requestId,"token无效")))
		return
	}
	if umap["mac"].(string) != claims["mac"].(string) {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(requestId,"token无效,请先登录")))
		return
	}
	tcpClient.SetIsLogin(true)

	userId:=claims["uid"].(string)
	util.Clients[userId]=tcpClient

	tcpClient.Write(base.Struct2Json(R.TcpOK(requestId)))
}

func chat(tcpClient *network.TcpClient,umap map[string]interface{})  {
	requestId:=umap["requestId"].(string)
	loginStruct:=LoginStruct{}
	if err := mapstructure.Decode(umap, &loginStruct); err != nil {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(requestId,"请求参数错误")))
	}
}

func move(tcpClient *network.TcpClient,umap map[string]interface{})  {

}