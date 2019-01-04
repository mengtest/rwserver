package dao

import (
	"../../rbstruct/user"
	"../../rbwork/base"
	"../../rbwork/db"
	"strconv"
)

func GetRolesByUserId(userId string) []user.RoleInfo {
	lUserId,_:=strconv.ParseInt(userId, 10, 64)
	roles:=[]user.RoleInfo{}
	sqlc:="SELECT * FROM tb_role WHERE nDeleted=0 AND lUserId=? ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sqlc," Param:",lUserId)
	err:=db.DB.Select(&roles,sqlc,lUserId)
	base.CheckErr(err)
	return roles
}

func GetRoleByRoleId(roleId string) user.RoleInfo {
	lRoleId,_:=strconv.ParseInt(roleId, 10, 64)
	role:=user.RoleInfo{}
	sqlc:="SELECT * FROM tb_role WHERE nDeleted=0 AND lId=? ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sqlc," Param:",lRoleId)
	err:=db.DB.Get(&role,sqlc,lRoleId)
	base.CheckErr(err)
	return role
}