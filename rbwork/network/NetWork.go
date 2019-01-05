package network

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
	"net"
	"strings"
	"time"
)

type TcpClient struct {
	ip        string        //客户端IP
	mac       string        //客户端mac地址
	conn      net.Conn      //客户端连接
	reader    *bufio.Reader //客户端输入读取缓冲区
	isLogin   bool          //是否通过登录授权访问
	timestamp int64         //上次心跳检测收到返回的时间戳（秒）
	userId    string        //用户ID
	roleId    string        //角色ID
}

func NewTcpClient(conn net.Conn) *TcpClient {
	ip := strings.Split(conn.RemoteAddr().String(), ":")[0] //获取客户端IP
	return &TcpClient{ip: ip, conn: conn, reader: bufio.NewReader(conn), isLogin: false, timestamp: time.Now().Unix()}
}

func (c *TcpClient) LocalAddr() string {
	return c.conn.LocalAddr().String()
}

func (c *TcpClient) RemoteAddr() string {
	return c.conn.RemoteAddr().String()
}

func (c *TcpClient) Close() error {
	return c.conn.Close()
}

func (c *TcpClient) GetIP() string {
	return c.ip
}

func (c *TcpClient) SetIsLogin(b bool) {
	c.isLogin = b
}

func (c *TcpClient) GetIsLogin() bool {
	return c.isLogin
}

func (c *TcpClient) SetTime(t int64) {
	c.timestamp = t
}

func (c *TcpClient) GetTime() int64 {
	return c.timestamp
}

func (c *TcpClient) SetUserId(userId string) {
	c.userId = userId
}

func (c *TcpClient) GetUserId() string {
	return c.userId
}

func (c *TcpClient) SetRoleId(roleId string) {
	c.roleId = roleId
}

func (c *TcpClient) GetRoleId() string {
	return c.roleId
}

func (c *TcpClient) SetMac(mac string) {
	c.mac = mac
}

func (c *TcpClient) GetMac() string {
	return c.mac
}

//不懂这里字节的话，建议先看下字节原理
func (c *TcpClient) Write(message string) (int, error) {
	// 读取消息的长度
	//tcp一次传输一般2个字节长度足够，所以这里用int16，如果不够可采用4字节的int32或者更高
	//根据debug显示int16下分配了两位uint8字节，最大可达255*256+255=65535，所以应付绝大多数场景是足够的
	var length = int16(len(message))
	if uint16(length) > 65535 {
		return 0, errors.New("byte length exceeded the maximum limit of 65535")
	}
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
	lengthByte, err := c.reader.Peek(2) //读取int16长度字节
	if err != nil {
		return "", err
	}

	lengthBuff := bytes.NewBuffer(lengthByte)

	var length int16
	err = binary.Read(lengthBuff, binary.BigEndian, &length)
	if err != nil {
		return "", err
	}
	// 字节超过了最大传输限制，返回错误
	if uint16(length) > 65535 {
		return "", errors.New("byte length exceeded the maximum limit of 65535")
	}

	//2字节长度处理逻辑
	ln := int16(c.reader.Buffered())
	if ln < length+2 {
		//这里有个问题，如果客户端只输入了报文长度，没有输入报文，则会出现这种请求
		return "", errors.New("could not read body info,please check client request")
	}
	pack := make([]byte, int(2+length))
	_, err = c.reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[2:]), nil

}
