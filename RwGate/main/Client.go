package main

import (
	"fmt"
	"net"
	"encoding/binary"
	"time"
	rw "../util"
)


func main() {
	conn, err := net.Dial("tcp", "localhost:1024")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	go Sender(conn)

	for {

		time.Sleep(1 * 1e9)

	}


}


func Sender(conn net.Conn) {

	//for i := 0; i < 100; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		rw.Log("发送报文",words)
		var headSize int
		var headBytes = make([]byte, 2)
		content := []byte(words)
		headSize = len(content)
		binary.BigEndian.PutUint16(headBytes, uint16(headSize))
		conn.Write(headBytes)
		conn.Write(content)

	//}

}