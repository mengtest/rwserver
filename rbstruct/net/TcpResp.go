package net

type Hurt struct {
	Value        int      `json:"value"`      //伤害值
	AttackType   int      `json:"attackType"` //攻击类型 1物理 2法术
	EffectType   int      `json:"effectType"` //0普通 1会心 2躲避 3要害
	SkillName    string   `json:"skillName"`  //技能名字
}

func NewHurt(v int,at int,et int,sn string) Hurt {
	return Hurt{Value:v,AttackType:at,EffectType:et,SkillName:sn}
}