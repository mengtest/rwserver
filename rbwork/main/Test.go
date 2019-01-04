package main

import (
	"../base"
	"fmt"
)

type Test struct {
	lId int
}

var TestMap =make(map[string]*Test)

func main() {
	// 3des 测试
	encMsg:=base.DesEncode("18758295232")
	fmt.Println(string(encMsg))
	decMsg,_:=base.DesDecode(encMsg)
	fmt.Println(string(decMsg))
}

