package handle

import (
	R "../../rbstruct/base"
	"../../rbwork/base"
	"../../rbwork/network"
	"reflect"
	"../service"
)

func HandleMsg(tcpClient *network.TcpClient,msg string)  {
	umap,err:=base.Json2map(msg)
	if err !=nil {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("非法报文")))
		return
	}
	cmd:=umap["cmd"]
	requestId:=umap["requestId"]
	if cmd ==nil || reflect.TypeOf(cmd).String() !="string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("无效请求")))
		return
	}
	if requestId ==nil ||  reflect.TypeOf(requestId).String() !="string"{
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("无效ID")))
		return
	}

	service:=&service.Service{}
	sv := reflect.ValueOf(&service).Elem()

	params := make([]reflect.Value,2)
	params[0] = reflect.ValueOf(tcpClient)
	params[1] = reflect.ValueOf(umap)

    //被调用方法名必须要大写,否则会抛异常
	m:=sv.MethodByName(cmd.(string))
	if m.IsValid() {
		m.Call(params)
	}else{
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(cmd.(string),requestId.(string),"无效请求")))
	}

}
