package network

import (
	"github.com/pkg/errors"
	"net"
	"bufio"
	"bytes"
	"encoding/binary"
	"strings"
    rwuser "../../RwStruct/user"
)

type TcpClient struct {
	ip  string            //客户端IP
	conn net.Conn         //客户端连接
	reader *bufio.Reader  //客户端输入读取缓冲区
	user rwuser.UserInfo  //用户
}

func NewTcpClient(conn net.Conn) *TcpClient {
	ip:=strings.Split(conn.RemoteAddr().String(),":")[0] //获取客户端IP
	return &TcpClient{ip: ip, conn: conn, reader: bufio.NewReader(conn)}
}

func (c *TcpClient) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *TcpClient) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *TcpClient) Close() error {
	return c.conn.Close()
}

func (c *TcpClient) GetIP() string  {
	return c.ip
}

func (c *TcpClient) Write(message string) (int, error) {
	// 读取消息的长度
	var length  = int32(len(message))
	var pkg = new(bytes.Buffer)

	//报文长度写入头
	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return 0, err
	}

	err = binary.Write(pkg, binary.BigEndian, []byte(message))
	if err != nil {
		return 0, err
	}

	n, err := c.conn.Write(pkg.Bytes())
	if err != nil {
		return 0, err
	}
	return n, nil
}

/**
 * 读取tcp请求报文，解决半包粘包
 * 先读取前4个字节得到报文长度，然后读取
 */
func (c *TcpClient) Read() (string, error) {
	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 个字节的数据，
	// 该操作不会将数据读出，只是引用，引用的数据在下一次读取操作之
	// 前是有效的。如果切片长度小于 n，则返回一个错误信息说明原因。
	// 如果 n 大于缓存的总大小，则返回 ErrBufferFull。
	lengthByte, err := c.reader.Peek(4)
	if err != nil {
		return "", err
	}

	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err = binary.Read(lengthBuff, binary.BigEndian, &length)
	if err != nil {
		return "", err
	}
	ln:=int32(c.reader.Buffered())
	if ln < length+4 {
		//这里有个问题，如果客户端只输入了报文长度，没有输入报文，则会出现这种请求
		return "", errors.New("could not read body info,please check client request")
	}

	pack := make([]byte, int(4+length))
	_, err = c.reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}

