package service

import (
	R "../../rbstruct/base"
	"../../rbwork/network"
	"../../rbstruct/user"
	"../../rbwork/constant"
	"../../rbwork/base"
	"../util"
	"time"
	"../../rbwork/redis"
	"strconv"
)

//定义Service机构体，反射调用其方法
type Service struct {
}

//ping 心跳时间更新
func (s *Service) Ping(tcpClient *network.TcpClient,msg string)  {
	tcpClient.SetTime(time.Now().Unix())
}

//同步自己的信息给周围玩家，package内部调用
func SyncPlayerToAroundPlayers(currRoleId string,role user.RoleInfo,buff *user.RoleBuff)  {
	//---- 获取周围角色ID
	var roleIds []string
	for i := role.NChunkX - 1; i <= role.NChunkX+1; i++ {
		for j := role.NChunkY - 1; j <= role.NChunkY+1; j++ {
			rIds := redis.Client.SMembers(constant.MapChunk + role.StrMapName + ":" + strconv.Itoa(i) + "#" + strconv.Itoa(j))
			roleIds = append(roleIds, rIds.Val()...)
		}
	}

	player := user.RespRole{}
	player.LId = role.LId
	player.StrName = role.StrName
	player.StrTitle = role.StrTitle
	player.NSex = role.NSex
	player.NLevel =role.NLevel
	player.NHP = role.NHP
	player.NMP = role.NMP
	player.NOccId = role.NOccId
	player.StrOccName = role.StrOccName
	player.FPosX = role.FPosX
	player.FPosY = role.FPosY
	player.FPosZ = role.FPosZ
	player.FDirX = role.FDirX
	player.FDirY = role.FDirY
	player.FDirZ = role.FDirZ
	player.Action= role.Action
	//产生新的buff时才向附近玩家推送新buff
	if buff != nil {
		player.Buffs=append(player.Buffs,*buff)
	}
     //Sync 表示同步消息命令 0代表立即执行
	for _, roleId := range roleIds {
		if roleId != "" && roleId !=currRoleId {
			client := util.Clients.Get(roleId)
			client.Write(base.Struct2Json(R.TcpOK("Sync", "0").SetData(player).OutLog()))
		}
	}
}