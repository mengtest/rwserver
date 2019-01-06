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
	TargetId int64 `json:"targetId"`
	TargetType int `json:"targetType"`
	SkillCode string `json:"skillCode"`
}

type ChatReq struct {
	Req
	ChatType int `json:"chatType"`
	ToRoleId int64 `json:"toRoleId"`
	Msg string `json:"msg"`
}