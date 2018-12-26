###tcp c/c++测试代码,服务端对半包粘包已处理

####网络编程中,TCP/IP统一采用大端方式传送数据，所以大端方式也称之为网络字节序

1.启动TQGate的GateServer  
2.运行测试代码
```
bool SendMsg(char* data)
{
	int len = strlen(data);

	char pBuf[1024];
	strcpy(pBuf, data);

	uint16 ulen= len;
	ulen = htons(ulen);//转网络传输字节序
	char buf[1024];
	memcpy(buf, &ulen, 2);  //前两位字节指定报文长度
	memcpy(buf+2, &pBuf, len);

	int outsize = send(m_sockClient, buf, len+2, 0);
	if (outsize <= 0) {
		Destroy();
		return false;
	}	
	return true;
}
```