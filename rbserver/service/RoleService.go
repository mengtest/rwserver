package service

import (
	R "../../rbstruct/base"
	"../../rbstruct/user"
	"../../rbwork/base"
	"../../rbwork/constant"
	"../../rbwork/network"
	"../../rbwork/redis"
	"../util"
	"reflect"
	"strconv"
)

//角色升级
func (s *Service) Upgrade(tcpClient *network.TcpClient, umap map[string]interface{}) {

}

//角色移动
func (s *Service) Move(tcpClient *network.TcpClient, umap map[string]interface{}) {
	mapName := umap["mapName"]
	chunkX := umap["chunkX"]
	chunkY := umap["chunkY"]
	px := umap["px"]
	py := umap["py"]
	pz := umap["pz"]
	dx := umap["dx"]
	dy := umap["dy"]
	dz := umap["dz"]
	if mapName == nil || reflect.TypeOf(mapName).String() != "string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("参数错误")))
		return
	}
	if chunkX == nil || reflect.TypeOf(chunkX).String() != "string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("参数错误")))
		return
	}
	if chunkX == nil || reflect.TypeOf(chunkX).String() != "string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("参数错误")))
		return
	}
	if chunkY == nil || reflect.TypeOf(chunkY).String() != "string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("参数错误")))
		return
	}
	if px == nil || reflect.TypeOf(px).String() != "string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("参数错误")))
		return
	}
	if py == nil || reflect.TypeOf(py).String() != "string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("参数错误")))
		return
	}
	if pz == nil || reflect.TypeOf(pz).String() != "string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("参数错误")))
		return
	}
	if dx == nil || reflect.TypeOf(dx).String() != "string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("参数错误")))
		return
	}
	if dy == nil || reflect.TypeOf(dy).String() != "string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("参数错误")))
		return
	}
	if dz == nil || reflect.TypeOf(dz).String() != "string" {
		tcpClient.Write(base.Struct2Json(R.ErrorMsg("参数错误")))
		return
	}
	nChunkX,_:=strconv.Atoi(chunkX.(string))
	nChunkY,_:=strconv.Atoi(chunkY.(string))

	role := tcpClient.GetRole()
    //如果所处地图和地块发生变化,角色属性变化
	if  mapName != role.StrMapName || nChunkX != role.NChunkX || nChunkY != role.NChunkY{
		redis.Client.SRem(constant.MapChunk +role.StrMapName+":"+ strconv.Itoa(role.NChunkX) + "#" + strconv.Itoa(role.NChunkY),tcpClient.GetRoleId())
		role.NChunkX=nChunkX
		role.NChunkY=nChunkY
		role.StrMapName=mapName.(string)
		redis.Client.SAdd(constant.MapChunk+role.StrMapName+":"+strconv.Itoa(role.NChunkX)+"#"+strconv.Itoa(role.NChunkY),tcpClient.GetRoleId())
	}
}

//角色攻击敌人
func (s *Service) Attack(tcpClient *network.TcpClient, umap map[string]interface{}) {

}

//角色接任务
func (s *Service) AcceptTask(tcpClient *network.TcpClient, umap map[string]interface{}) {

}

//角色放弃任务
func (s *Service) AbandonTask(tcpClient *network.TcpClient, umap map[string]interface{}) {

}

//角色完成任务
func (s *Service) FinishTask(tcpClient *network.TcpClient, umap map[string]interface{}) {

}

//角色获取物品（装备、物品）
func (s *Service) GetGoods(tcpClient *network.TcpClient, umap map[string]interface{}) {

}

//角色丢弃物品
func (s *Service) DiscardGoods(tcpClient *network.TcpClient, umap map[string]interface{}) {

}

//获取周围玩家列表
func (s *Service) GetAroundPlayers(tcpClient *network.TcpClient, umap map[string]interface{}) {
	requestId := umap["requestId"].(string)
	if !tcpClient.GetIsLogin() {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("GetAroundPlayers", requestId, "未授权，请先登录")))
		return
	}
	if tcpClient.GetRoleId() == "" {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("GetAroundPlayers", requestId, "未选择角色")))
		return
	}

	role := tcpClient.GetRole()
	//---- 获取周围角色ID
	var roleIds []string
	for i := role.NChunkX - 1; i <= role.NChunkX+1; i++ {
		for j := role.NChunkY - 1; j <= role.NChunkY+1; j++ {
			rIds:= redis.Client.SMembers(constant.MapChunk +role.StrMapName+":"+ strconv.Itoa(i) + "#" + strconv.Itoa(j))
			roleIds = append(roleIds, rIds.Val() ...)
		}
	}
	//---- 获取这些角色信息
	var players []user.RoleInfo
	for _, roleId := range roleIds {
		if roleId != "" {
			client:=util.Clients.Get(roleId)
			players = append(players, *client.GetRole())
		}
	}
	tcpClient.Write(base.Struct2Json(R.TcpOK("GetAroundPlayers", requestId).SetData(players).OutLog()))

}
