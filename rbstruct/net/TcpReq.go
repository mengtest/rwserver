package net

type Req struct {
	Cmd       string `json:"cmd"`
	RequestId string `json:"requestId"`
}

//-- 继承Req
type LoginReq struct {
	Req
	Token string `json:"token"`
	Mac   string `json:"mac"`
}
