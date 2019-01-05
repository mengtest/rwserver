package service

import (
	R "../../rbstruct/base"
	"../../rbstruct/user"
	"../../rbwork/base"
	"../../rbwork/constant"
	"../../rbwork/network"
	"../../rbwork/redis"
	"../util"
	"strconv"
)

//角色升级
func (s *Service) Upgrade(tcpClient *network.TcpClient, umap map[string]interface{}) {

}

//角色移动
func (s *Service) Move(tcpClient *network.TcpClient, umap map[string]interface{}) {

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
			rIds, _ := redis.Smembers(constant.MAP_CHUNK + role.StrMapName + ":" + strconv.Itoa(i) + ":" + strconv.Itoa(j))
			roleIds = append(roleIds, rIds...)
		}
	}
	//---- 获取这些角色信息
	var players []user.RoleInfo
	for _, roleId := range roleIds {
		if roleId != "" {
			base.LogInfo("roleId==" + roleId)
			players = append(players, *util.Clients.Get(roleId).GetRole())
		}
	}
	tcpClient.Write(base.Struct2Json(R.TcpOK("GetAroundPlayers", requestId).SetData(players).OutLog()))

}
