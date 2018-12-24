package base

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"encoding/json"
	"reflect"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func GetParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
//打印结构体
func LogStruct(v interface{})  {
	if v==nil {
       return
	}
	b, err := json.Marshal(v)
	if err == nil {
       LogInfo(string(b))
	}
}

//利用反射判断interface是否为空
func IsNil(v interface{}) bool {
	vi := reflect.ValueOf(&v)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}