package main

import (
	"../base"
	"fmt"
)
func main() {
	// 3des 测试
	encMsg:=base.DesEncode([]byte("18758295232"))
	fmt.Println(string(encMsg))
	decMsg,_:=base.DesDecode(encMsg)
	fmt.Println(string(decMsg))
}
