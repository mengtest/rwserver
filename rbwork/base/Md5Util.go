package base

import (
	"crypto/md5"
)

func GetMd5(str string) string{
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return string(cipherStr)
}

