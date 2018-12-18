package main

import (
	"strconv"
	"../util"
)

func main()  {
	s:="85ba"
	i, err :=strconv.ParseInt(s,16,32)
	if err != nil {
		panic(err)
	}
	j:=strconv.IntSize;
	util.Log("i=",i,"j=",j)
}