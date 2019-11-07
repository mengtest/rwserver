package main

import (
	"../../rbwork/base"
	"../../rbwork/db"
	"../../rbwork/network"
	"../../rbwork/redis"
	"../dao"
	"../handle"
	"../service"
	Gloat "../util"
	"net"
	"os"
	"time"
)

func init() {
	base.Init(base.GetCurrentDirectory(), "GameServer.log")
	db.Init("root:123456@tcp(127.0.0.1:3306)/tianqi?charset=utf8")
	redis.InitRedis("127.0.0.1:6379", "Wyz123!@#")
	service.InitLevel()
}

func main() {

	//建立socket，监听端口  第一步:绑定端口
	netListen, err := net.Listen("tcp", "localhost:9010")
	CheckError(err)
	//defer延迟关闭改资源，以免引起内存泄漏
	defer netListen.Close()
	base.LogInfo("game sever start")

	go runHeartbeat()

	go runSyncRoleInfoToDB()

	for {
		conn, err := netListen.Accept()
		if err != nil {
			base.LogError(err)
			continue //出错退出当前循环
		}
		//实例化TcpClient,方便进行统一管理
		client := network.NewTcpClient(conn)
		base.LogInfo(client.RemoteAddr(), "===>tcp connect success")
		//将连接加入全局map 此时该连接没有经过认证，待心跳检测超时无返回时则中断该连接，并从map清除
		//登录认证的用户重新设置用户ID为主键

		Gloat.Clients.Set(client.RemoteAddr(), client)
		//使用协程处理并发请求
		go handleConnection(client)

	}
}

//处理连接
func handleConnection(tcpClient *network.TcpClient) {
	//这里不主动关闭连接，由系统心跳检测去统一管理

	for {
		message, err := tcpClient.Read()
		if err != nil {
			base.LogError(err)
			//监听到客户端退出，关闭连接
			tcpClient.Close()
			Gloat.Clients.Delete(tcpClient.RemoteAddr())
			Gloat.Clients.Delete(tcpClient.GetStrRoleId())
			return
		}
		//输出收到的日志信息
		base.LogInfo("accept:", message)
		go handle.HandleMsg(tcpClient, message)
	}

}

//处理error
func CheckError(err error) {
	if err != nil {
		base.LogError(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//心跳检测
func runHeartbeat() {
	//每5秒执行一次检测
	tick := time.NewTicker(time.Second * time.Duration(5))
	for {
		<-tick.C
		//base.LogInfo("开始发送心跳包")
		for _, tcpClient := range Gloat.Clients.GetMap() {
			timeb := time.Now().Unix() - tcpClient.GetTime() //计算秒
			if timeb > 40 {
				//40s内未收到心跳返回,剔除用户
				tcpClient.Close()
				Gloat.Clients.Delete(tcpClient.RemoteAddr())
				Gloat.Clients.Delete(tcpClient.GetStrRoleId())
				base.LogInfo("IP->"+tcpClient.RemoteAddr(), "roleId->"+tcpClient.GetStrRoleId(), "超过40秒未收到心跳返回，已断开连接")
			}
			//tcpClient.Write("{\"cmd\":\"ping\",\"requestId\":\"ping\"}")
		}
	}
	defer tick.Stop()
}

//同步角色信息到DB
func runSyncRoleInfoToDB() {
	//每10秒执行一次
	tick := time.NewTicker(time.Second * time.Duration(10))
	for {
		<-tick.C
		base.LogInfo("开始同步角色到数据库")
		dao.SyncRoleToDB()
	}
	defer tick.Stop()
}
