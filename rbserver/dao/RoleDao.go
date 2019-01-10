package dao

import (
	"../../rbstruct/user"
	"../../rbwork/base"
	"../../rbwork/db"
	"strconv"
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

func GetRoleSkillByRoleId(roleId int64) []user.RoleSkill {
	roleSkill:=[]user.RoleSkill{}
	sqlc:="SELECT * FROM tb_role_skill WHERE lRoleId=? "
	base.LogInfo("SQL:",sqlc," Param:",roleId)
	err:=db.DB.Select(&roleSkill,sqlc,roleId)
	base.CheckErr(err)
	return roleSkill
}

func GetRoleSQL(role user.RoleInfo) string {
	sql:="UPDATE tb_role SET "
	sql+="strName='"+role.StrName+"',"
	sql+="strTitle='"+role.StrTitle+"',"
	sql+="WHERE lId="+strconv.FormatInt(role.LId,10)
	return sql
}

func ExecSQL(sql string)  {
	db.DB.MustBegin()
	db.DB.MustExec(sql)

}