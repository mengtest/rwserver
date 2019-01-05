package main

import (
	"../../rbwork/base"
	"../../rbwork/network"
	"fmt"
	"net"
	"time"
)

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "localhost:9010")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")
	time.Sleep(1 * time.Microsecond)
	go sendMessage(conn)
	go Recv(conn)
	var msg string
	fmt.Scanln(&msg)

}

func sendMessage(conn *net.TCPConn) {

	//time.Sleep(1 * time.Microsecond)
	//content:="{\"code\":\"asdjkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkweqweopqweqopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopop\",\"msg\":\"weweeweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee\"}"
	content := "{\"cmd\":\"Login\",\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDc5OTA1MzUsImlhdCI6MTU0NjY5NDUzNSwiaXNzIjoidHEuaXVvb24uY29tIiwibWFjIjoiIiwidWlkIjoiXHUwMDAxIn0.wOK-LIk-gXlhLlRPZ1H8mFamkiCzgZaZBE1s7URReNE\",\"requestId\":\"1\",\"mac\":\"\"}"
	b, _ := base.EncodeHead2Byte(string(content))
	conn.Write(b)

	content = "{\"cmd\":\"LoginRole\",\"roleId\":\"1\",\"requestId\":\"1\",\"mac\":\"\"}"
	b, _ = base.EncodeHead2Byte(string(content))
	conn.Write(b)

	time.Sleep(1 * time.Second)
	content = "{\"cmd\":\"GetAroundPlayers\",\"roleId\":\"1\",\"requestId\":\"1\",\"mac\":\"\"}"
	b, _ = base.EncodeHead2Byte(string(content))
	conn.Write(b)


}

func Recv(conn *net.TCPConn) {
	tcpClient := network.NewTcpClient(conn)
    for {
		message, err := tcpClient.Read()
		if err != nil {
			base.LogError(err)
			return
		}
		//输出收到的日志信息
		base.LogInfo("收到返回:", message)
	}
}
