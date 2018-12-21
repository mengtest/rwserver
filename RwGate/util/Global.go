package util

import (
	"net"
	"log"
	"strconv"
	"bytes"
	"encoding/binary"
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

var Count int


//log输出
func Log(v ...interface{}) {
	log.Println(v...)
}

//16进制转10进制数字
func HexToTen(str string)uint16{
	i, err :=strconv.ParseInt(str,16,32)
	if err != nil {
		Log(err)
	}
	return uint16(i)
}


func IntToByte(num int) []byte {
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, num)
	return buffer.Bytes()
}
