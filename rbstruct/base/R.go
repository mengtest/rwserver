package base

import ( TQ "../../rbwork/base")

//定义返回结构体  字段后面跟`json:"fieldName"` 可以设定显示别名  注意字段首字母必须大写，否则返回前端的时候不显示
type R struct {
	Code int   `json:"code"`
	Msg string  `json:"msg"`
	Data interface{} `json:"data,omitempty"`
	RequestId string `json:"requestId,omitempty"`
	Cmd string `json:"cmd,omitempty"`
}

func OK() R{
	return R{Code: 0, Msg: "处理成功"}
}

func OkMsg(message string) R{
	return R{Code: 0, Msg: message}
}

func OkCodeMsg(code int,message string) R{
	return R{Code: code, Msg: message}
}

func Error() R{
	return R{Code: -1, Msg: "系统异常"}
}

func ErrorMsg(message string) R{
	return R{Code: -1, Msg: message }
}

func ErrorCodeMsg(code int,message string) R{
	return R{Code: code, Msg: message}
}

func (r R) SetData(data interface{}) R{
	r.Data=data
	return r
}

func (r R) OutLog() R{
	//异步输出该结构体
    go TQ.LogStruct("backMsg:",r)
	return r
}

func TcpOK(cmd,requestId string) R{
	return R{Code: 0, Msg: "处理成功",RequestId:requestId,Cmd:cmd}
}

func TcpOkMsg(cmd,requestId ,message string) R{
	return R{Code: 0, Msg: message,RequestId:requestId,Cmd:cmd}
}

func TcpError(cmd,requestId string) R{
	return R{Code: -1, Msg: "系统异常",RequestId:requestId,Cmd:cmd}
}

func TcpErrorMsg(cmd,requestId ,message string) R{
	return R{Code: -1, Msg: message ,RequestId:requestId,Cmd:cmd}
}