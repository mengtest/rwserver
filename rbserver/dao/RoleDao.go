package dao

import (
	"../../rbstruct/user"
	"../../rbwork/base"
	"../../rbwork/db"
)

func GetRolesByUserId(userId int64) []user.RoleInfo {
	roles:=[]user.RoleInfo{}
	sqlc:="SELECT * FROM tb_role WHERE nDeleted=0 AND lUserId=? ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sqlc," Param:",userId)
	err:=db.DB.Select(&roles,sqlc,userId)
	base.CheckErr(err)
	return roles
}

func GetRoleByRoleId(roleId int64) user.RoleInfo {
	role:=user.RoleInfo{}
	sqlc:="SELECT * FROM tb_role WHERE nDeleted=0 AND lId=? ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sqlc," Param:",roleId)
	err:=db.DB.Get(&role,sqlc,roleId)
	base.CheckErr(err)
	return role
}