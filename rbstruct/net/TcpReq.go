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

type AttackReq struct {
	Req
	TargetId int64    `json:"targetId"`       //目标ID （角色ID，怪物ID）
	TargetType int    `json:"targetType"`     //目标类型0-角色 1-怪物
	SkillId int64     `json:"skillId"`       //技能ID
}

type ChatReq struct {
	Req
	NChannel       int       `json:"nChannel"`
	LFromRoleId    int64     `json:"lFromRoleId"`
	LToRoleId      int64     `json:"lToRoleId"`
	StrMsg         string    `json:"strMsg"`
}

type RealeaseSkill struct {
	Req
	SkillId    int64    `json:"skillId"`
}