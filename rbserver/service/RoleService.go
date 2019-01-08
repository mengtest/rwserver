package service

import (
	R "../../rbstruct/base"
	"../../rbstruct/net"
	"../../rbstruct/user"
	"../../rbwork/base"
	"../../rbwork/constant"
	"../../rbwork/network"
	"../../rbwork/redis"
	"../util"
	"strconv"
)

//角色升级
func (s *Service) Upgrade(tcpClient *network.TcpClient, msg string) {

}

//角色移动
func (s *Service) Move(tcpClient *network.TcpClient, msg string) {
	req := &net.MoveReq{}
	base.Json2Struct(msg, req)

	role := tcpClient.GetRole()
	//如果所处地图和地块发生变化,角色属性变化
	if req.MapName != role.StrMapName || req.ChunkX != role.NChunkX || req.ChunkY != role.NChunkY {
		redis.Client.SRem(constant.MapChunk+role.StrMapName+":"+strconv.Itoa(role.NChunkX)+"#"+strconv.Itoa(role.NChunkY), tcpClient.GetRoleId())
		role.NChunkX = req.ChunkX
		role.NChunkY = req.ChunkY
		role.StrMapName = req.MapName
		redis.Client.SAdd(constant.MapChunk+role.StrMapName+":"+strconv.Itoa(role.NChunkX)+"#"+strconv.Itoa(role.NChunkY), tcpClient.GetRoleId())
	}
	role.FDirX = req.Dx
	role.FDirY = req.Dy
	role.FDirZ = req.Dz
	role.FPosX = req.Px
	role.FPosY = req.Py
	role.FPosZ = req.Pz
	//同步消息
	SyncPlayerToAroundPlayers(tcpClient.GetStrRoleId(),*tcpClient.GetRole(),nil)
}

//角色攻击敌人
func (s *Service) Attack(tcpClient *network.TcpClient, msg string) {
	req := &net.AttackReq{}
	base.Json2Struct(msg, req)
	if req.TargetId <= 0 {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(req.Cmd, req.RequestId, "请选择目标")))
		return
	}
	role:=tcpClient.GetRole()
	targetClient:=util.Clients.Get(strconv.FormatInt(req.TargetId,10))
	skill:=user.RoleSkill{}
	for _,sk:= range role.Skills  {
		if sk.LSkillId == req.LSkillId {
			skill=sk
			break
		}
	}
	//计算伤害
	v:=0
	if req.LSkillId == 0 {
		//计算普攻伤害
		v=base.RandInt(role.NMinAD,role.NMaxAD)*5/4-targetClient.GetRole().NPhyDef*2/3
		if targetClient.GetRole().NHP <= v {
			targetClient.GetRole().NHP=0
			targetClient.GetRole().Action="die"
		}else{
			targetClient.GetRole().NHP-=v
		}
	} else {
		if skill.NAttackType==1{
			v=base.RandInt(role.NMinAD,role.NMaxAD)*5/4+skill.NSkillValue-targetClient.GetRole().NPhyDef*2/3
		}else if skill.NAttackType==2{
			v=base.RandInt(role.NMinAP,role.NMaxAP)*5/4+skill.NSkillValue-targetClient.GetRole().NMagDef*2/3
		}
		if targetClient.GetRole().NHP <= v {
			targetClient.GetRole().NHP=0
			targetClient.GetRole().Action="die"
		}else{
			targetClient.GetRole().NHP-=v
		}
	}
	role.Action="attack_"+skill.StrSkillCode
	//向被攻击者推送伤害
	targetClient.Write(base.Struct2Json(R.TcpOK(R.Hurt_Own,"0").SetData(nil)))
	//向攻击发起者推送伤害
	tcpClient.Write(base.Struct2Json(R.TcpOK(R.Hurt_Target, "0").SetData(nil)))
    //同步被攻击者信息给附近玩家
	SyncPlayerToAroundPlayers(targetClient.GetStrRoleId(),*targetClient.GetRole(),nil)
	//同步攻击者信息给附近玩家
	SyncPlayerToAroundPlayers(tcpClient.GetStrRoleId(),*tcpClient.GetRole(),nil)
}

//释放技能
func (s *Service) ReleaseSkill(tcpClient *network.TcpClient, msg string) {

}

//角色接任务
func (s *Service) AcceptTask(tcpClient *network.TcpClient, msg string) {

}

//角色放弃任务
func (s *Service) AbandonTask(tcpClient *network.TcpClient, msg string) {

}

//角色完成任务
func (s *Service) FinishTask(tcpClient *network.TcpClient, msg string) {

}

//角色获取物品（装备、物品）
func (s *Service) GetGoods(tcpClient *network.TcpClient, msg string) {

}

//角色丢弃物品
func (s *Service) DiscardGoods(tcpClient *network.TcpClient, msg string) {

}

//获取周围玩家列表
func (s *Service) GetAroundPlayers(tcpClient *network.TcpClient, msg string) {
	req := &net.Req{}
	base.Json2Struct(msg, req)
	if !tcpClient.GetIsLogin() {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(req.Cmd, req.RequestId, "未授权，请先登录")))
		return
	}
	if tcpClient.GetRoleId() <= 0 {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg(req.Cmd, req.RequestId, "未选择角色")))
		return
	}

	role := tcpClient.GetRole()
	//---- 获取周围角色ID
	var roleIds []string
	for i := role.NChunkX - 1; i <= role.NChunkX+1; i++ {
		for j := role.NChunkY - 1; j <= role.NChunkY+1; j++ {
			rIds := redis.Client.SMembers(constant.MapChunk + role.StrMapName + ":" + strconv.Itoa(i) + "#" + strconv.Itoa(j))
			roleIds = append(roleIds, rIds.Val()...)
		}
	}
	//---- 获取这些角色信息
	var players []user.RespRole
	for _, roleId := range roleIds {
		if roleId != "" {
			client := util.Clients.Get(roleId)

			player := user.RespRole{}
			player.LId = client.GetRole().LId
			player.StrName = client.GetRole().StrName
			player.StrTitle = client.GetRole().StrTitle
			player.NSex = client.GetRole().NSex
			player.NLevel = client.GetRole().NLevel
			player.NHP = client.GetRole().NHP
			player.NMP = client.GetRole().NMP
			player.NMaxHP = client.GetRole().NMaxHP
			player.NMaxMP = client.GetRole().NMaxMP
			player.NOccId = client.GetRole().NOccId
			player.StrOccName = client.GetRole().StrOccName

			players = append(players, player)
		}
	}
	tcpClient.Write(base.Struct2Json(R.TcpOK(req.Cmd, req.RequestId).SetData(players).OutLog()))

}
