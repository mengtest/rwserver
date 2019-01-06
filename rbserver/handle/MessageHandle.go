package handle

import (
	R "../../rbstruct/base"
	"../../rbstruct/net"
	"../../rbwork/base"
	"../../rbwork/network"
	"reflect"
	"../service"
)

func HandleMsg(tcpClient *network.TcpClient,msg string)  {
	req:=&net.Req{}
	base.Json2Struct(msg,req)

	if req.Cmd=="" || req.RequestId=="" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("无效请求")))
		return
	}

	service:=&service.Service{}
	sv := reflect.ValueOf(&service).Elem()

	params := make([]reflect.Value,2)
	params[0] = reflect.ValueOf(tcpClient)
	params[1] = reflect.ValueOf(msg)

    //被调用方法名必须要大写,否则会抛异常
	m:=sv.MethodByName(req.Cmd)
	if m.IsValid() {
		m.Call(params)
	}else{
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(req.Cmd,req.RequestId,"无效请求")))
	}

}
