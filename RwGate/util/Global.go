package util

import "net"

//定义全局map存储客户端连接clients
var ClientMap map[uint32]net.Conn

var Count int