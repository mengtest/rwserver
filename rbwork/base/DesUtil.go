package base

import (
	"crypto/des"
	"crypto/cipher"
	"bytes"
	"encoding/base64"
)

const DesKey  =   "wuhantianqi20181227rbctx"

//解密
func DesDecode(encMsg string)(string,error){
	crypted, err := base64.StdEncoding.DecodeString(encMsg)
	if err != nil {
	  return "",err
	}
	key := []byte(DesKey)
	//获取block块
	block,_ :=des.NewTripleDESCipher(key)
	//创建切片
	context := make([]byte,len(crypted))
	//设置解密方式
	blockMode := cipher.NewCBCDecrypter(block,key[:8])
	//解密密文到数组
	blockMode.CryptBlocks(context,crypted)
	//去补码
	context = PKCSUnPadding(context)
	return string(context),nil
}

//去补码
func PKCSUnPadding(origData []byte)[]byte{
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:length-unpadding]
}
//加密
func DesEncode(decMsg string)string{
	key := []byte(DesKey)
	origData:=[]byte(decMsg)
	//获取block块
	block,_ :=des.NewTripleDESCipher(key)
	//补码
	origData = PKCSPadding(origData, block.BlockSize())
	//设置加密方式为 3DES  使用3条56位的密钥对数据进行三次加密
	blockMode := cipher.NewCBCEncrypter(block,key[:8])
	//创建明文长度的数组
	crypted := make([]byte,len(origData))
	//加密明文
	blockMode.CryptBlocks(crypted,origData)
	return base64.StdEncoding.EncodeToString(crypted) //返回base64编码
}
//补码
func PKCSPadding(origData []byte,blockSize int)[]byte{
	//计算需要补几位数
	padding := blockSize-len(origData)%blockSize
	//在切片后面追加char数量的byte(char)
	padtext := bytes.Repeat([]byte{byte(padding)},padding)
	return append(origData,padtext...)
}
