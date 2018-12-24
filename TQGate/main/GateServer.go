package main

import (
	TQ "../../TQBase/base"
	TQNet "../../TQBase/network"
	"net"
	"os"
)

func main() {
	TQ.Init(TQ.GetCurrentDirectory(),"GateServer.log")
	//建立socket，监听端口  第一步:绑定端口
	netListen, err := net.Listen("tcp", "localhost:9090")
	CheckError(err)
	//defer延迟关闭改资源，以免引起内存泄漏
	defer netListen.Close()

	TQ.LogInfo("gate sever start")
	for {
		conn, err := netListen.Accept()  //第二步:获取连接
		if err != nil {
			TQ.LogError(err)
			continue  //出错退出当前循环
		}
        //实例化TcpClient,方便进行统一管理
		client := TQNet.NewTcpClient(conn)
		TQ.LogInfo(client.GetIP(), "===>tcp connect success")
		//将连接加入全局map

		//使用协程处理并发请求
		go handleConnection(client)

	}
}
//处理连接
func handleConnection(tcpClient *TQNet.TcpClient) {
	//这里服务端先不着急关闭连接，由系统心跳检测去统一管理
	//defer tcpClient.Close()

	for {
		message, err := tcpClient.Read()
		if err != nil {
			TQ.LogError(err)
			return
		}
        //输出收到的日志信息
		TQ.LogInfo("accept:", message)

	}

}



//处理error
func CheckError(err error) {
	if err != nil {
		TQ.LogError(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}


