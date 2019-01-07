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
	NExp         int            `db:"nExp" json:"nExp"`
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
}

//角色技能
type RoleSkill struct {
	LId          int64          `db:"lId" json:"lId"`
	LRoleId      int64          `db:"lRoleId" json:"lRoleId"`
	LSkillId     int64          `db:"lSkillId" json:"lSkillId"`
	StrSkillName string         `db:"strSkillName" json:"strSkillName"`
	NLevel       int            `db:"nLevel" json:"nLevel"`
	NSkillValue  int            `db:"nSkillValue" json:"nSkillValue"`
	NSkillType   int            `db:"nSkillType" json:"nSkillType"`
	StrDesc      string         `db:"strDesc" json:"strDesc"`
}

//角色buff
type RoleBuff struct {
	LSkillId     int64          `json:"lSkillId"` //技能ID，由什么技能触发的BUF
	NLevel       int64          `json:"nLevel"`   //技能等级
	StrProp      string         `json:"strProp"`  //影响角色的哪项属性
	StrDesc      string         `json:"strDesc"`  //状态描述
	NSeconds     int            `json:"nSeconds"` //剩余持续时间（秒）
	NType        int            `json:"nType"`    //1增益型buf -1减益型buf
	NValue       int            `json:"nType、"`  //增益值或减益值
}

//周围角色
type AroundRole struct {
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
}
