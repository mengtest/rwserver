package user

import "database/sql"

//定义角色结构体
type RoleInfo struct {
	LId          int64          `db:"lId" json:"lId"`
	LUserId      int64          `db:"lUserId" json:"lUserId"`
	StrName      string         `db:"strName" json:"strName"`
	StrTitle     string         `db:"strTitle" json:"strTitle"`
	NSex         int            `db:"nSex" json:"nSex"`
	NLevel       int            `db:"nLevel" json:"nLevel"`
	NCurtExp     int64          `db:"nCurtExp" json:"nCurtExp"`
	NExp         int64          `db:"nExp" json:"nExp"`
	NHP          int            `db:"nHP" json:"nHP"`                //当前血气值
	NMP          int            `db:"nMP" json:"nMP"`
	NMaxHP       int            `db:"nMaxHP" json:"nMaxHP"`          //最大血气值
	NMaxMP       int            `db:"nMaxMP" json:"nMaxMP"`

	NMinAP       int            `db:"nMinAP" json:"nMinAP"`
	NMinAD       int            `db:"nMinAD" json:"nMinAD"`
	NMaxAP       int            `db:"nMaxAP" json:"nMaxAP"`
	NMaxAD       int            `db:"nMaxAD" json:"nMaxAD"`
	NPhyDef      int            `db:"nPhyDef" json:"nPhyDef"`
	NMagDef      int            `db:"nMagDef" json:"nMagDef"`
	NDodge       int            `db:"nDodge" json:"nDodge"`
	NCrit        int            `db:"nCrit" json:"nCrit"`
	NHit         int            `db:"nHit" json:"nHit"`
	NCon         int            `db:"nCon" json:"nCon"`
	NDex         int            `db:"nDex" json:"nDex"`
	NStr         int            `db:"nStr" json:"nStr"`
	NAvoid       int            `db:"nAvoid" json:"nAvoid"`
	NSp          int            `db:"nSp" json:"nSp"`
	FPosX        float64        `db:"fPosX" json:"fPosX"`
	FPosY        float64        `db:"fPosY" json:"fPosY"`
	FPosZ        float64        `db:"fPosZ" json:"fPosZ"`
	FDirX        float64        `db:"fDirX" json:"fDirX"`
	FDirY        float64        `db:"fDirY" json:"fDirY"`
	FDirZ        float64        `db:"fDirZ" json:"fDirZ"`
	StrMapName   string         `db:"strMapName" json:"strMapName"`
	NChunkX      int            `db:"nChunkX" json:"nChunkX"`
	NChunkY      int            `db:"nChunkY" json:"nChunkY"`
	NOccId       int            `db:"nOccId" json:"nOccId"`
	StrOccName   string         `db:"strOccName" json:"strOccName"`
	NDeleted     int            `db:"nDeleted" json:"-"`
	DtDeleteDate sql.NullString `db:"dtDeleteDate" json:"-"`
	DtUpdateTime sql.NullString `db:"dtUpdateTime" json:"-"`
	DtCreateTime string         `db:"dtCreateTime" json:"-"`
	Skills       []RoleSkill    `json:"skills"`                //角色技能
	Buffs        []RoleBuff     `json:"buffs"`                //角色状态
	Action       string         `json:"action"`
}

//所有技能都是法术，未使用技能为普通物理攻击
//角色技能
type RoleSkill struct {
	LId          int64          `db:"lId" json:"lId"`
	LRoleId      int64          `db:"lRoleId" json:"lRoleId"`
	LSkillId     int64          `db:"lSkillId" json:"lSkillId"`
	StrSkillCode string         `db:"strSkillCode" json:"strSkillCode"`  //技能编码（action动作编码）
	StrSkillName string         `db:"strSkillName" json:"strSkillName"`
	NSkillLevel  int            `db:"nSkillLevel" json:"nSkillLevel"`
	NSkillValue  int            `db:"nSkillValue" json:"nSkillValue"`
	NSkillType   int            `db:"nSkillType" json:"nSkillType"`   //1攻击 2状态Buff
	NAttackType  int            `db:"nAttackType" json:"nAttackType"` //1物理加成 2法术加成
	StrDesc      string         `db:"strDesc" json:"strDesc"`         //技能介绍
	StrEffectDesc string        `db:"strEffectDesc" json:"strEffectDesc"`  //效果描述
	NCastTime    int            `db:"nCastTime" json:"nCastTime"`     //施法时间
	NDuration    int            `db:"nDuration" json:"nDuration"`     //持续时间
}

//角色buff
type RoleBuff struct {
	LSkillId     int64          `json:"lSkillId"`       //技能ID，由什么技能触发的BUF
	StrSkillName      string    `json:"strSkillName"`   //strName
	NLevel       int64          `json:"nLevel"`         //技能等级
	StrProp      string         `json:"strProp"`        //影响角色的哪项属性
	StrEffectDesc      string   `json:"strEffectDesc"`  //状态描述
	NType        int            `json:"nType"`          //1增益型buf -1减益型buf
	NValue       int            `json:"nValue"`         //增益值或减益值
	NDuration    int            `json:"nDuration"`     //持续时间
}


type RespRole struct {
	LId          int64          `db:"lId" json:"lId"`
	StrName      string         `db:"strName" json:"strName"`
	StrTitle     string         `db:"strTitle" json:"strTitle"`
	NSex         int            `db:"nSex" json:"nSex"`
	NLevel       int            `db:"nLevel" json:"nLevel"`
	NHP          int            `db:"nHP" json:"nHP"`
	NMP          int            `db:"nMP" json:"nMP"`
	NMaxHP       int            `db:"nMaxHP" json:"nMaxHP"`
	NMaxMP       int            `db:"nMaxMP" json:"nMaxMP"`
	NOccId       int            `db:"nOccId" json:"nOccId"`
	StrOccName   string         `db:"strOccName" json:"strOccName"`
	FPosX        float64        `db:"fPosX" json:"fPosX"`
	FPosY        float64        `db:"fPosY" json:"fPosY"`
	FPosZ        float64        `db:"fPosZ" json:"fPosZ"`
	FDirX        float64        `db:"fDirX" json:"fDirX"`
	FDirY        float64        `db:"fDirY" json:"fDirY"`
	FDirZ        float64        `db:"fDirZ" json:"fDirZ"`
	Buffs        []RoleBuff     `json:"buffs"`                //角色状态
	Action       string         `json:"action"`               //角色动作 idle(idle0 idle1 idle2)、walk、run、dying(死亡中，执行死亡动作)、die(倒地)、releaseSkill(技能释放)、hurt（执行受伤动作）
}
