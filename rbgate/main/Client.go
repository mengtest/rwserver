package main

import (
	"fmt"
	"net"
	RW "../../rbwork/base"
	"time"
)


func main() {

	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "localhost:9090")

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
		content:="{\"code\":\"asdjkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkweqweopqweqopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopop\",\"msg\":\"weweeweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee\"}"
		//content:="{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},{3456712},"
		b, _ := RW.EncodeHead2Byte(string(content))
		conn.Write(b)
	}
}



