package handle

import (
	"../../rbwork/base"
	"../../rbwork/network"
	R "../../rbstruct/base"
	"github.com/goinggo/mapstructure"
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
	cmd:=umap["cmd"].(string)
	if cmd =="" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("无效请求")))
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
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("无效请求")))
		return
	}

}

func login(tcpClient *network.TcpClient,umap map[string]interface{})  {
	strToken:=umap["token"].(string)
	claims,err :=base.DecodeToken(strToken)
	if err != nil {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("token无效")))
		return
	}
	if umap["mac"].(string) != claims["mac"].(string) {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("token无效,请先登录")))
		return
	}
	tcpClient.Write(base.Struct2Json(R.OK()))
}

func chat(tcpClient *network.TcpClient,umap map[string]interface{})  {
	loginStruct:=LoginStruct{}
	if err := mapstructure.Decode(umap, &loginStruct); err != nil {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("请求参数错误")))
	}
}

func move(tcpClient *network.TcpClient,umap map[string]interface{})  {

}