package main

import (
	"fmt"
	"net"
	"../../rbwork/base"
	"../../rbwork/network"
	"strconv"
	"time"
)


func main() {
	for i:=0;i<50;i++ {
		base.LogInfo(base.GenId())
	}

	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "localhost:9010")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")
	time.Sleep(1 * time.Microsecond)
	go sendMessage(conn)



	var msg string
	fmt.Scanln(&msg)

}


func sendMessage(conn *net.TCPConn) {
	for i:=0;i<10 ;i++ {
		//time.Sleep(1 * time.Microsecond)
		//content:="{\"code\":\"asdjkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkweqweopqweqopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopop\",\"msg\":\"weweeweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee\"}"
		content:="{\"cmd\":\"Login\",\"token\":\"2323123\",\"requestId\":\""+strconv.Itoa(i)+"\"}"
		b, _ := base.EncodeHead2Byte(string(content))
		conn.Write(b)

        Recv(conn)
	}
}

func Recv(conn *net.TCPConn)  {
	tcpClient := network.NewTcpClient(conn)

		message, err := tcpClient.Read()
		if err != nil {
			base.LogError(err)
			return
		}
		//输出收到的日志信息
		base.LogInfo("收到返回:", message)

}

