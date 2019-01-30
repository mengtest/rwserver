package net

type Req struct {
	Cmd       string `json:"cmd"`
	RequestId string `json:"requestId"`
}

//-- 继承Req
type LoginReq struct {
	Req
	Token string `json:"token"`
	Mac   string `json:"mac"`
}

type LoginRoleReq struct {
	Req
	RoleId int64 `json:"roleId"`
}

type MoveReq struct {
	Req
	MapName string `json:"mapName"`
	ChunkX  int `json:"chunkX"`
	ChunkY  int `json:"chunkY"`
	Px  float64 `json:"px"`
	Py  float64 `json:"py"`
	Pz  float64 `json:"pz"`
	Dx  float64 `json:"dx"`
	Dy  float64 `json:"dy"`
	Dz  float64 `json:"dz"`
}
//攻击消息
type AttackReq struct {
	Req
	TargetId int64    `json:"targetId"`       //目标ID （角色ID，怪物ID）
	TargetType int    `json:"targetType"`     //目标类型0-角色 1-怪物
	SkillId int64     `json:"skillId"`       //技能ID
}
//发送消息
type ChatReq struct {
	Req
	NChannel       int       `json:"nChannel"`
	LFromRoleId    int64     `json:"lFromRoleId"`
	LToRoleId      int64     `json:"lToRoleId"`
	StrMsg         string    `json:"strMsg"`
}
//升级
type UpgradeReq struct {
	Req
}
//-- 增加阅历
type IncreaseExpReq struct {
	Req
	NExp           int64       `json:"nExp"`
}

type CreateRoleReq struct {
	Req
	StrName        string      `json:"strName"`
	NSex           string      `json:"nSex"`
	NOccId         int         `json:"nOccId"`
	StrOccName     string      `json:"strOccName"`
}
