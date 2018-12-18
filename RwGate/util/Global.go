package util

import (
	"net"
	"log"
	"strconv"
)

//定义全局map存储客户端连接clients
var ClientMap map[uint32]net.Conn

var Count int


//log输出
func Log(v ...interface{}) {
	log.Println(v...)
}

//16进制转10进制数字
func toUint64(str string)uint64{
	i, err :=strconv.ParseInt(str,16,64)
	if err != nil {
		Log(err)
	}
	return uint64(i)
}