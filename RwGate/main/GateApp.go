package main

import (
	"fmt"
	"net"
	"os"
	rw "../util"
	"bufio"
)

func main() {

	//建立socket，监听端口  第一步:绑定端口
	netListen, err := net.Listen("tcp", "localhost:1024")
	//netListen, err := net.Listen("tcp", "127.0.0.1:9800")
	CheckError(err)
	//defer延迟关闭改资源，以免引起内存泄漏
	defer netListen.Close()

	rw.Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()  //第二步:获取连接
		if err != nil {
			continue  //出错退出当前一次循环
		}

		rw.Log(conn.RemoteAddr().String(), " tcp connect success")
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
		message, err := rw.Decode(reader)
		if err != nil {
			return
		}

		rw.Log(err,"accept:[",conn.RemoteAddr().String() ,"]:" , string(message))

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
