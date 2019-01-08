package service

import (
	"time"
	"../../rbstruct/user"
	"../util"
	)

//定时器管理
var BuffTimerManager= make(map[string]*time.Timer)

//buff 持续时间及持续结束后计算
/**
 * @param roleId 角色ID
 * @param second 持续时间(秒)
 * @param buff   角色buff
 */
func BuffCalculation(roleId string,second int,buff user.RoleBuff)  {
	t:=BuffTimerManager[roleId+buff.StrProp]
	if t ==nil {
		t = time.NewTimer(time.Second * time.Duration(second))
		BuffTimerManager[roleId+buff.StrProp]=t
	}else{
		t.Reset(time.Second * time.Duration(second))
	}
	select {
	    case <-t.C:
			switch buff.StrProp {
			case "nMaxHP":
				util.Clients.Get(roleId).GetRole().NMaxHP-=buff.NValue*buff.NType
				break
			case "nMaxMP":
				util.Clients.Get(roleId).GetRole().NMaxMP-=buff.NValue*buff.NType
				break
			case "nMinAP":
				util.Clients.Get(roleId).GetRole().NMinAP-=buff.NValue*buff.NType
				break
			case "nMinAD":
				util.Clients.Get(roleId).GetRole().NMaxAD-=buff.NValue*buff.NType
				break
			case "nMaxAP":
				util.Clients.Get(roleId).GetRole().NMaxAP-=buff.NValue*buff.NType
				break
			case "nMaxAD":
				util.Clients.Get(roleId).GetRole().NMaxAD-=buff.NValue*buff.NType
				break
			case "nPhyDef":
				util.Clients.Get(roleId).GetRole().NPhyDef-=buff.NValue*buff.NType
				break
			case "nMagDef":
				util.Clients.Get(roleId).GetRole().NMagDef-=buff.NValue*buff.NType
				break
			case "nDodge":
				util.Clients.Get(roleId).GetRole().NDodge-=buff.NValue*buff.NType
				break
			case "nCrit":
				util.Clients.Get(roleId).GetRole().NCrit-=buff.NValue*buff.NType
				break
			case "nHit":
				util.Clients.Get(roleId).GetRole().NHit-=buff.NValue*buff.NType
				break
			case "nCon":
				util.Clients.Get(roleId).GetRole().NCon-=buff.NValue*buff.NType
				break
			case "nDex":
				util.Clients.Get(roleId).GetRole().NDex-=buff.NValue*buff.NType
				break
			case "nStr":
				util.Clients.Get(roleId).GetRole().NStr-=buff.NValue*buff.NType
				break
			case "nAvoid":
				util.Clients.Get(roleId).GetRole().NAvoid-=buff.NValue*buff.NType
				break
			case "nSp":
				util.Clients.Get(roleId).GetRole().NSp-=buff.NValue*buff.NType
				break
			default:
				break
			}
			//清除对应时间管理器
			delete(BuffTimerManager,roleId+buff.StrProp)
	        //清除角色的buff
			DeleteRoleBuff(roleId,buff)
			//改变后的状态同步给周围玩家
			SyncPlayerToAroundPlayers(roleId,*util.Clients.Get(roleId).GetRole(),nil)
	}
	defer t.Stop()
}

func DeleteRoleBuff(roleId string,buff user.RoleBuff)  {
	buffs:=util.Clients.Get(roleId).GetRole().Buffs
	for i,b:= range buffs  {
       if b.LSkillId==buff.LSkillId{
		   buffs = append(buffs[:i], buffs[i+1:]...)
		   break
	   }
	}
	util.Clients.Get(roleId).GetRole().Buffs=buffs
}