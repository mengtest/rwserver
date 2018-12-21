package util

import (
	"net"
)

//用户
type UserClient struct {
	conn net.Conn   //套接字
	strIP    string //请求IP
	lUserId  uint32 //用户ID
	nChannelNo int  //渠道号
}

//定义全局存储在线用户，键名用户ID
var Users = make(map[uint32]UserClient)
//定义全局map存储客户端连接,键名ip地址
var Conns = make(map[string]net.Conn)
