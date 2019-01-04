package user

//定义角色结构体
type RoleInfo struct {
	LId int64 `db:"lId" json:"lId"`
	StrName string `db:"strName" json:"strName"`
	StrTitle string `db:"strTitle" json:"strTitle"`
	NSex string `db:"nSex" json:"nSex"`
	NLevel string `db:"nLevel" json:"nLevel"`
	NExp int `db:"nExp" json:"nExp"`
	NHP string `db:"nHP" json:"nHP"`
	NMP  string `db:"nMP" json:"nMP"`
	NMinAP string `db:"nMinAP" json:"nMinAP"`
	NMinAD string `db:"nMinAD" json:"nMinAD"`
	NMaxAP int `db:"nMaxAP" json:"nMaxAP"`
	NMaxAD int `db:"nMaxAD" json:"nMaxAD"`
	NPhyDef int `db:"nPhyDef" json:"nPhyDef"`
	NMagDef int `db:"nMagDef" json:"nMagDef"`
	NDodge int `db:"nDodge" json:"nDodge"`
	NCrit int `db:"nCrit" json:"nCrit"`
	NHit int `db:"nHit" json:"nHit"`
	NCon int `db:"nCon" json:"nCon"`
	NDex int `db:"nDex" json:"nDex"`
	NStr int `db:"nStr" json:"nStr"`
	NAvoid int `db:"nAvoid" json:"nAvoid"`
	NSp int `db:"nSp" json:"nSp"`
	FPosX int `db:"fPosX" json:"fPosX"`
	FPosY int `db:"fPosY" json:"fPosY"`
	FPosZ int `db:"fPosZ" json:"fPosZ"`
	FDirX int `db:"fDirX" json:"fDirX"`
	FDirY int `db:"fDirY" json:"fDirY"`
	FDirZ int `db:"fDirZ" json:"fDirZ"`
	StrMapName int `db:"strMapName" json:"strMapName"`
	NDeleted int `db:"nDeleted" json:"-"`
	DtDeleteDate int `db:"dtDeleteDate" json:"-"`
	DtUpdateTime int `db:"dtUpdateTime" json:"-"`
	DtCreateTime int `db:"dtCreateTime" json:"-"`
}
