package util

import (
	"net"
	rwnet "../../rbwork/network"
)



//定义全局存储在线用户，键名用户ID
var Clients = make(map[string]*rwnet.TcpClient)
//定义全局map存储客户端连接,键名ip地址
var Conns = make(map[string]net.Conn)
