package main

import (
	"fmt"
	"net"
	"os"
	RwUtil "../../RwBase/base"
	//RwNet "../../RwBase/network"
	"bufio"
)

func main() {

	//建立socket，监听端口  第一步:绑定端口
	netListen, err := net.Listen("tcp", "localhost:1024")
	CheckError(err)
	//defer延迟关闭改资源，以免引起内存泄漏
	defer netListen.Close()

	RwUtil.Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()  //第二步:获取连接
		if err != nil {
			continue  //出错退出当前一次循环
		}

		RwUtil.Log(conn.RemoteAddr().String(), " tcp connect success")
		//将连接加入全局map

		//handleConnection(conn)  //正常连接就处理
		//这句代码的前面加上一个 go，就可以让服务器并发处理不同的Client发来的请求
		go handleConnection(conn) //使用goroutine来处理用户的请求

	}
}
//处理连接
func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := RwUtil.Decode(reader)
		if err != nil {
			RwUtil.Log("err","端户断开连接")
			return
		}
		if message == "" {
			continue
		}

		RwUtil.Log("accept:[",conn.RemoteAddr().String() ,"]:" , message)

		//b, err := rw.Encode(conn.RemoteAddr().String() + ":" + string(message))
		//if err != nil {
		//	continue
		//}
		//conn.Write(b)

	}

}



//处理error
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}


