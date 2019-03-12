package service

import (
	"../dao"
	R "../../rbstruct/base"
	"../../rbstruct/net"
	"../../rbwork/base"
	"../../rbwork/network"
	Gloal "../util"
)

//加载等级配置
func InitLevel()  {
	levels:=dao.GetLevelConfig()
	for _,level:= range levels  {
		Gloal.LevelConfig[level.NLevel]=level.NExp
	}
}

//获取等级配置
func (s *Service) GetLevelConfig(tcpClient *network.TcpClient, msg string) {
	req := &net.Req{}
	base.Json2Struct(msg, req)
	tcpClient.Write(base.Struct2Json(R.TcpOK(req.Cmd, req.RequestId).SetData(Gloal.LevelConfig).OutLog()))
}