package base

import (
	"encoding/json"
	"github.com/goinggo/mapstructure"
	"github.com/satori/go.uuid"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"
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
func LogStruct(prefix string, v interface{}) {
	if v == nil {
		return
	}
	b, err := json.Marshal(v)
	if err == nil {
		LogInfo(prefix, string(b))
	}
}

func CheckErr(err error) bool {
	if err != nil {
		LogError(err)
		return true
	}
	return false
}

func Json2map(jsonstr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonstr), &m)
	return m, err
}

func Struct2Json(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func Json2Struct(jsonstr string, s *interface{}) error {
	err := json.Unmarshal([]byte(jsonstr), s)
	return err
}

func MapToStruct(m map[string]interface{}, s *interface{}) error {
	err := mapstructure.Decode(m, s)
	return err
}

//GetIncreaseID 并发环境下生成一个增长的id,按需设置局部变量或者全局变量
func GetIncreaseID(ID *uint64) uint64 {
	var n, v uint64
	for {
		v = atomic.LoadUint64(ID)
		n = v + 1
		if atomic.CompareAndSwapUint64(ID, v, n) {
			break
		}
	}
	rand.Seed(time.Now().UnixNano())
	return n
}

func GenId() string {
	u, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return strings.Replace(u.String(), "-", "", -1)
}
