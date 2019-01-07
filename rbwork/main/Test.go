package main

import (
	"../base"
	"fmt"
	"strconv"
	"time"
)

type Test struct {
	lId int
}

var TestMap = make(map[string]*Test)

func main() {
	// 3des 测试
	encMsg := base.DesEncode("18758295232")
	fmt.Println(string(encMsg))
	decMsg, _ := base.DesDecode(encMsg)
	fmt.Println(string(decMsg))

	//测试生成token
	var a int64
	a = 1
	fmt.Println(base.CreateToken(strconv.FormatInt(a, 10), ""))

	//测试定时惹怒
	t := time.NewTimer(time.Second * time.Duration(3))
	select {
	case <-t.C:
		fmt.Println("执行定时任务")
	}
	defer t.Stop()

	//测试
}
