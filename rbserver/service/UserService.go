package service

import (
	R "../../rbstruct/base"
	"../../rbwork/base"
	"../../rbwork/db"
	"../../rbwork/network"
	"../util"
)

//定义角色结构体
type Role struct {
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

//登录授权校验
func (s *Service) Login(tcpClient *network.TcpClient,umap map[string]interface{})  {
	strToken:=umap["token"].(string)
	requestId:=umap["requestId"].(string)
	claims,err :=base.DecodeToken(strToken)
	if err != nil {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("Login",requestId,"token无效")))
		return
	}
	if umap["mac"].(string) != claims["mac"].(string) {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("Login",requestId,"token无效,请先登录")))
		return
	}
	userId:=claims["uid"].(string)
	tcpClient.SetIsLogin(true)
	tcpClient.SetUserId(userId)
	tcpClient.SetMac(umap["mac"].(string))

	tcpClient.Write(base.Struct2Json(R.TcpOK("Login",requestId)))
}


//获取角色列表
func (s *Service) GetRoles(tcpClient *network.TcpClient,umap map[string]interface{})  {
	requestId:=umap["requestId"].(string)
	if !tcpClient.GetIsLogin() {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("Login",requestId,"未授权，请先登录")))
		return
	}

	roles:=[]Role{}
	sqlc:="SELECT * FROM tb_role WHERE nDeleted=0 AND lUserId=? ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sqlc," Param:",tcpClient.GetUserId())
	err:=db.DB.Select(&roles,sqlc,tcpClient.GetUserId())
	base.CheckErr(err)

	tcpClient.Write(base.Struct2Json(R.TcpOK("GetRoles",requestId).SetData(roles).OutLog()))
}

//选择角色进入
func (s *Service) LoginRole(tcpClient *network.TcpClient,umap map[string]interface{})  {
	requestId:=umap["requestId"].(string)
	roleId:=umap["roleId"].(string)
	//load role info
	if !tcpClient.GetIsLogin() {
		tcpClient.Write(base.Struct2Json(R.TcpErrorMsg("Login",requestId,"未授权，请先登录")))
		return
	}

	tcpClient.SetRoleId(roleId)

	util.Clients.Delete(tcpClient.GetIP()) //清除游客模式连接
	util.Clients.Set(roleId,tcpClient)     //设置用户ID为主键
	tcpClient.Write(base.Struct2Json(R.TcpOK("LoginRole",requestId)))
}


