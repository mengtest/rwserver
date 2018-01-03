package service

import (
	"../../rbwork/network"
	"time"
)

//定义Service机构体，反射调用其方法
type Service struct {
}

//ping 心跳时间更新
func (s *Service) Ping(tcpClient *network.TcpClient,umap map[string]interface{})  {
	tcpClient.SetTime(time.Now().Unix())
}