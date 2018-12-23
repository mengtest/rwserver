package base

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

//头部4字节 4字节占用空间多（资源浪费，没达到这个量级不建议使用）
func EncodeHead4Byte(message string) ([]byte, error) {
	var length  = int32(len(message))
	var pkg  = new(bytes.Buffer)

	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.BigEndian, []byte(message))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

//头部4字节
func DecodeHead4Byte(reader *bufio.Reader) (string, error) {
	lengthByte, err := reader.Peek(4)
	if err != nil {
		return "", err
	}


	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err = binary.Read(lengthBuff, binary.BigEndian, &length)
	if err != nil {
		return "", err
	}
	ln:=int32(reader.Buffered())
	if ln < length+4 {
		return "", err
	}

	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}

//头部两字节
func EncodeHead2Byte(message string) ([]byte, error) {
	var length  = int16(len(message))
	var pkg  = new(bytes.Buffer)

	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.BigEndian, []byte(message))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

//头部两字节
func DecodeHead2Byte(reader *bufio.Reader) (string, error) {
	lengthByte, err := reader.Peek(2)
	if err != nil {
		return "", err
	}
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int16
	err = binary.Read(lengthBuff, binary.BigEndian, &length)
	if err != nil {
		return "", err
	}
	ln:=int16(reader.Buffered())
	if ln < length+4 {
		return "", err
	}

	pack := make([]byte, int(2+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[2:]), nil
}