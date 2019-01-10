package net

//定义伤害结构体
type Hurt struct {
	Value        int      `json:"value"`      //伤害值
	AttackType   int      `json:"attackType"` //攻击类型 1物理 2法术
	EffectType   int      `json:"effectType"` //0普通 1会心 2躲避 3要害
	SkillName    string   `json:"skillName"`  //技能名字
}

func NewHurt(v int,at int,et int,sn string) Hurt {
	return Hurt{Value:v,AttackType:at,EffectType:et,SkillName:sn}
}

//定义聊天结构体
type Chat struct {
	LFromRoleId      int64   `json:"lFromRoleId"`     //消息发起人ID      系统ID=0
	StrFromRoleName  string  `json:"strFromRoleName"` //消息发起人角色名  系统名称=系统
	LToRoleId        int64   `json:"lToRoleId"`       //消息接收人ID      未指定时，传0
	StrToRoleName    string  `json:"strToRoleName"`   //消息接收人角色名   未指定时，传空
	StrMsg           string  `json:"strMsg"`          //消息内容
	NChannel         int     `json:"nChannel"`        //发送渠道
}

func NewChat(fid int64,frn string,tid int64,trn string,msg string,c int) Chat {
	return Chat{LFromRoleId:fid,StrFromRoleName:frn,LToRoleId:tid,StrToRoleName:trn,StrMsg:msg,NChannel:c}
}