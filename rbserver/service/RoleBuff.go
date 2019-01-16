package service

import (
	"../../rbstruct/user"
	"../util"
	"time"
)

//定时器管理
var BuffTimerManager = make(map[string]*time.Timer)

//buff 持续时间及持续结束后计算
/**
 * @param roleId 角色ID
 * @param second 持续时间(秒)
 * @param buff   角色buff
 */
func BuffCalculation(roleId string, second int, buff user.RoleBuff) {
	t := BuffTimerManager[roleId+buff.StrProp]
	if t == nil {
		t = time.NewTimer(time.Second * time.Duration(second))
		BuffTimerManager[roleId+buff.StrProp] = t
	} else {
		t.Reset(time.Second * time.Duration(second))
	}
	select {
	case <-t.C:
		role := util.Clients.Get(roleId).GetRole()
		v := buff.NValue * buff.NType
		switch buff.StrProp {
		case "nMaxHP":
			role.NTempMaxHP -= v
			role.NMaxHP -= v
			break
		case "nMaxMP":
			role.NTempMaxMP -= v
			role.NMaxMP -= v
			break
		case "nMinAP":
			role.NTempMinAP -= v
			role.NMinAP -= v
			break
		case "nMinAD":
			role.NTempMaxAD -= v
			role.NMaxAD -= v
			break
		case "nMaxAP":
			role.NTempMaxAP -= v
			role.NMaxAP -= v
			break
		case "nMaxAD":
			role.NTempMaxAD -= v
			role.NMaxAD -= v
			break
		case "nPhyDef":
			role.NTempPhyDef -= v
			role.NPhyDef -= v
			break
		case "nMagDef":
			role.NTempMagDef -= v
			role.NMagDef -= v
			break
		case "nDodge":
			role.NTempDodge -= v
			role.NDodge -= v
			break
		case "nCastValue":
			role.NTempCastValue -= v
			role.NCastValue -= v
			break
		case "nCrit":
			role.NTempCrit -= v
			role.NCrit -= v
			break
		case "nHit":
			role.NTempHit -= v
			role.NHit -= v
			break
		case "nCon":
			role.NTempCon -= v
			role.NCon -= v
			break
		case "nDex":
			role.NTempDex -= v
			role.NDex -= v
			break
		case "nStr":
			role.NTempStr -= v
			role.NStr -= v
			break
		case "nAvoid":
			role.NTempAvoid -= v
			role.NAvoid -= v
			break
		case "nSp":
			role.NTempSp -= v
			role.NSp -= v
			break
		default:
			break
		}
		//清除对应时间管理器
		delete(BuffTimerManager, roleId+buff.StrProp)
		//清除角色的buff
		DeleteRoleBuff(roleId, buff)
		//改变后的状态同步给周围玩家
		SyncPlayerToAroundPlayers(roleId, *util.Clients.Get(roleId).GetRole(), nil)
	}
	defer t.Stop()
}

func DeleteRoleBuff(roleId string, buff user.RoleBuff) {
	buffs := util.Clients.Get(roleId).GetRole().Buffs
	for i, b := range buffs {
		if b.LSkillId == buff.LSkillId {
			buffs = append(buffs[:i], buffs[i+1:]...)
			break
		}
	}
	util.Clients.Get(roleId).GetRole().Buffs = buffs
}
