package main

import (
	"fmt"
	"net"
	"os"
	RwUtil "../../RwBase/base"
	"strings"
	//RwNet "../../RwBase/network"
	"bufio"
)

func main() {
	RwUtil.Init(RwUtil.GetCurrentDirectory(),"server.log")
	//建立socket，监听端口  第一步:绑定端口
	netListen, err := net.Listen("tcp", "localhost:1024")
	CheckError(err)
	//defer延迟关闭改资源，以免引起内存泄漏
	defer netListen.Close()

	RwUtil.LogInfo("Waiting for clients")
	for {
		conn, err := netListen.Accept()  //第二步:获取连接
		if err != nil {
			continue  //出错退出当前一次循环
		}
		clientIP:=strings.Split(conn.RemoteAddr().String(),":")[0] //获取客户端IP
		RwUtil.LogInfo(clientIP, "===>tcp connect success")
		//将连接加入全局map

		//这句代码的前面加上一个 go，就可以让服务器并发处理不同的Client发来的请求
		go handleConnection(conn) //使用goroutine来处理用户的请求

	}
}
//处理连接
func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	count:=0
	for {
		message, err := RwUtil.Decode(reader)
		if err != nil {
			RwUtil.LogInfo("err:",err)
			return
		}
		if message == "" {
			//这里防止要cpu空转，导致cpu内存使用率飙升
			if count>5 {
				count+=1
				continue
			}else{
				//5次读不到数据，则返回
				return
			}
		}else{
			count=0
		}
        //输出收到的日志信息
		RwUtil.LogInfo("accept:", message)

	}

}



//处理error
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}


