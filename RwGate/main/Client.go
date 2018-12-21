package main

import (
	"fmt"
	"net"
	RW "../../RwBase/base"
	"time"
)


func main() {

	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:1024")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")
	for i:=0;i<100 ;i++ {
	 	sendMessage(conn)
	}
	var msg string
	fmt.Scanln(&msg)

}


func sendMessage(conn *net.TCPConn) {

		time.Sleep(1 * time.Microsecond)
		content:="{\"code\":\"asdjkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkweqweopqweqopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopopop\",\"msg\":\"weweeweeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee\"}"
		b, _ := RW.Encode(string(content))
		conn.Write(b)

}



