package base

import "log"

//log输出
func Log(v ...interface{}) {
	log.Println(v...)
}