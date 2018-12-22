package base

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(message string) ([]byte, error) {
	var length  = int32(len(message))
	var pkg  = new(bytes.Buffer)

	//if(length>9 && length<99){
	//	binary.Write(pkg, binary.BigEndian, 0)
	//}
	//if(length>99 && length<999){
	//	binary.Write(pkg, binary.BigEndian, 0)
	//	binary.Write(pkg, binary.BigEndian, 0)
	//}

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

func Decode(reader *bufio.Reader) (string, error) {
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