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
	req := &net.UpgradeReq{}
	base.Json2Struct(msg, req)
	role := tcpClient.GetRole()
	newLevel:=role.NLevel+1
	levelC:=util.LevelMap[newLevel]

	if newLevel<=60 && levelC !=nil && role.NCurtExp>=levelC.NExp {
		role.NLevel=newLevel
		role.NCurtExp=role.NCurtExp-levelC.NExp
		tcpClient.Write(base.Struct2Json(R.TcpOK(req.Cmd, req.RequestId)))
	}else{
		tcpClient.Write(base.Struct2Json(R.TcpError(req.Cmd, req.RequestId)))
	}
}

//角色移动
func (s *Service) Move(tcpClient *network.TcpClient, msg string) {
	req := &net.MoveReq{}
	base.Json2Struct(msg, req)

	role := tcpClient.GetRole()
	//如果所处地图和地块发生变化,角色属性变化
	if req.MapName != role.StrMapName || req.ChunkX != role.NChunkX || req.ChunkY != role.NChunkY {
		//--在该旧地块上移除角色ID
		redis.Client.SRem(constant.MapChunk+role.StrMapName+":"+strconv.Itoa(role.NChunkX)+"#"+strconv.Itoa(role.NChunkY), tcpClient.GetRoleId())
		role.NChunkX = req.ChunkX
		role.NChunkY = req.ChunkY
		role.StrMapName = req.MapName
		//--在新地块加入角色ID
		redis.Client.SAdd(constant.MapChunk+role.StrMapName+":"+strconv.Itoa(role.NChunkX)+"#"+strconv.Itoa(role.NChunkY), tcpClient.GetRoleId())
	}
	role.FDirX = req.Dx
	role.FDirY = req.Dy
	role.FDirZ = req.Dz
	role.FPosX = req.Px
	role.FPosY = req.Py
	role.FPosZ = req.Pz
	role.BChange=true
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
		if sk.LSkillId == req.SkillId {
			skill=sk
			break
		}
	}
	//如果是释放buff技能
	if skill.NSkillType==2{

		buff:=user.RoleBuff{}
		buff.LSkillId=skill.LSkillId
		buff.StrSkillName=skill.StrSkillName
		buff.StrEffectDesc=skill.StrEffectDesc
		buff.NDuration=skill.NDuration-1 //因为通信时间差，服务器预先减一秒
		buff.StrProp=skill.StrProp
		role.Action="attack_"+skill.StrSkillCode
		buff.NValue=skill.NSkillValue
		buff.StrImgPath=skill.StrImgPath

		if req.TargetId == role.LId {
			//目标为自己
			buff.NType=1 //增益型
			SyncPlayerToAroundPlayers(tcpClient.GetStrRoleId(),*tcpClient.GetRole(),&buff)
			//buff持续时间计算，结束后进行结算
			go BuffCalculation(tcpClient.GetStrRoleId(),skill.NDuration,buff)
		}else{
			//目标为敌人
			buff.NType=-1 //减益型
			SyncPlayerToAroundPlayers(targetClient.GetStrRoleId(),*tcpClient.GetRole(),&buff)
			//buff持续时间计算，结束后进行结算
			go BuffCalculation(targetClient.GetStrRoleId(),skill.NDuration,buff)
		}
		return
	}
	//-----------------------计算伤害start----------------------------
	v:=0
	attackType:=1
	if req.SkillId == 0 {
		//计算普攻伤害
		v=base.RandInt(role.NMinAD,role.NMaxAD)*5/4-targetClient.GetRole().NPhyDef*2/3
		if targetClient.GetRole().NHP <= v {
			targetClient.GetRole().NHP=0
			targetClient.GetRole().Action="die"
		}else{
			targetClient.GetRole().NHP-=v
			targetClient.GetRole().Action="hurt"
		}
		attackType=1
	} else {

		//使用法术技能
		if skill.NAttackType==1 && skill.NSkillType==1 {
			//物理加成
			v=base.RandInt(role.NMinAD,role.NMaxAD)*5/4+skill.NSkillValue-targetClient.GetRole().NPhyDef*2/3
		}else if skill.NAttackType==2 && skill.NSkillType==1{
			//法术加成
			v=base.RandInt(role.NMinAP,role.NMaxAP)*5/4+skill.NSkillValue-targetClient.GetRole().NMagDef*2/3
		}else if skill.NSkillType==2{
			//buff
		}
		if targetClient.GetRole().NHP <= v {
			targetClient.GetRole().NHP=0
			targetClient.GetRole().Action="die"
			targetClient.GetRole().BChange=true
		}else{
			targetClient.GetRole().NHP-=v
			targetClient.GetRole().Action="hurt"
			targetClient.GetRole().BChange=true
		}
		attackType=2
	}
	//-----------------------计算伤害end----------------------------
	role.Action="attack_"+skill.StrSkillCode

	h:=net.NewHurt(v,attackType,1,skill.StrSkillName)
	//向被攻击者推送伤害
	targetClient.Write(base.Struct2Json(R.TcpOK(R.Hurt_Own,"0").SetData(h)))
	//向攻击发起者推送伤害
	tcpClient.Write(base.Struct2Json(R.TcpOK(R.Hurt_Target, "0").SetData(h)))
    //同步被攻击者信息给附近玩家
	SyncPlayerToAroundPlayers(targetClient.GetStrRoleId(),*targetClient.GetRole(),nil)
	//同步攻击者信息给附近玩家
	SyncPlayerToAroundPlayers(tcpClient.GetStrRoleId(),*tcpClient.GetRole(),nil)
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

//增加经验
func (s *Service) IncreaseExp(tcpClient *network.TcpClient, msg string) {
	req := &net.IncreaseExpReq{}
	base.Json2Struct(msg, req)
	role:=tcpClient.GetRole()
	role.NExp=role.NExp+req.NExp
	tExp:=role.NCurtExp+req.NExp
	if role.NLevel <= 10 && util.LevelMap[role.NLevel].NExp < tExp {
		role.NCurtExp=tExp-util.LevelMap[role.NLevel].NExp
		//10以下自动升级
		role.NLevel=role.NLevel+1
		//各属性自动加1
		role.NSp=role.NSp+1        //法     1法=4法攻  + 4法力   +  2命中 + 2法防
		role.NStr=role.NStr+1      //力     1力=4物攻  + 2命中   +  2物防
		role.NDex=role.NDex+1      //敏     1敏=4施法  + 4会心   +  2闪避
		role.NAvoid=role.NAvoid+1  //避     1避=4闪避  + 4会防   +  2物防 + 2法防
		role.NCon=role.NCon+1      //体     1体=10生命 + 2物防
		//计算防御、攻击等值
		role.NHP=role.NHP + 10
		role.NPhyDef=role.NPhyDef + 2 + 2
		role.NMagDef=role.NMagDef + 2 + 2
		role.NMP=role.NMP + 4
		role.NCrit=role.NCrit + 4
		role.NHit=role.NHit + 2 + 2
		role.NMaxAD=role.NMaxAD+4
		role.NMinAD=role.NMinAD+2
		role.NMaxAP=role.NMaxAP+4
		role.NMinAP=role.NMinAP+2
		role.NCritDef=role.NCritDef+4
		role.NDodge=role.NDodge+4
	}else{
		role.NExp=role.NExp+req.NExp
		role.NCurtExp=role.NCurtExp+req.NExp
	}

}
