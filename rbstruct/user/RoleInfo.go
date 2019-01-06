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
	NHP          int            `db:"nHP" json:"nHP"`
	NMP          int            `db:"nMP" json:"nMP"`
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
	NDeleted     int            `db:"nDeleted" json:"-"`
	DtDeleteDate sql.NullString `db:"dtDeleteDate" json:"-"`
	DtUpdateTime sql.NullString `db:"dtUpdateTime" json:"-"`
	DtCreateTime string         `db:"dtCreateTime" json:"-"`
}
