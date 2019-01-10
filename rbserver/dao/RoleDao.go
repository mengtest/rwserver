package dao

import (
	"../../rbstruct/user"
	"../../rbwork/base"
	"../../rbwork/db"
	"../util"
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

//将角色数据同步到数据库
func SyncRoleToDB()  {
	tx := db.DB.MustBegin()
	for _, client := range util.Clients.GetMap() {
		r:=client.GetRole()
		//同步角色到数据库
		if client.GetRole().BChange {
			sql:="UPDATE tb_role SET "
			sql+="strName=?,"
			sql+="strTitle=?,"
			sql+="nLevel=?,"
			sql+="nCurtExp=?,"
			sql+="nExp=?,"
			sql+="nHP=?,"
			sql+="nMP=?,"
			sql+="nMinAP=?,"
			sql+="nMinAD=?,"
			sql+="nMaxAP=?,"
			sql+="nMaxAD=?,"
			sql+="nPhyDef=?,"
			sql+="nMagDef=?,"
			sql+="nDodge=?,"
			sql+="nCrit=?,"
			sql+="nHit=?,"
			sql+="nCon=?,"
			sql+="nDex=?,"
			sql+="nStr=?,"
			sql+="nAvoid=?,"
			sql+="nSp=?,"
			sql+="fPosX=?,"
			sql+="fPosY=?,"
			sql+="fPosZ=?,"
			sql+="fDirX=?,"
			sql+="fDirY=?,"
			sql+="fDirZ=?,"
			sql+="strMapName=?,"
			sql+="nChunkX=?,"
			sql+="nChunkY=?,"
			sql+="dtUpdateTime=NOW()"
			sql+="WHERE lId=?"
			tx.MustExec(sql,r.StrName,r.StrTitle,r.NLevel,r.NCurtExp,r.NExp,r.NHP-r.NTempHP,r.NMP-r.NTempMP,r.NMinAP-r.NTempMinAP,r.NMinAD-r.NTempMinAD,r.NMaxAP-r.NTempMaxAP,r.NMaxAD-r.NTempMaxAD,
				r.NPhyDef-r.NTempPhyDef,r.NMagDef-r.NTempMagDef,r.NDodge-r.NTempDodge,r.NCrit-r.NTempCrit,r.NHit-r.NTempHit,r.NCon-r.NTempCon,r.NDex-r.NTempCon,r.NStr-r.NTempStr,r.NAvoid-r.NTempAvoid,r.NSp-r.NTempSp,r.FPosX,r.FPosY,r.FPosZ,r.FDirX,r.FDirY,r.FDirZ,r.StrMapName,r.NChunkX,r.NChunkY,r.LId)
		}
		//同步角色技能到数据库
		if client.GetRole().BChangeSkill {
			for _, s := range client.GetRole().Skills {
				if s.BChange && s.ChangeType == 1 {
					skilSql:="INSERT INTO tb_role_skill(lRoleId,lSkillId,strSkillCode,strSkillName,nSkillLevel,nSkillValue,nSkillType,nAttackType,strDesc,nCastTime,nDuration,strEffectDesc) " +
						"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)"
					skilSql+=""
					tx.MustExec(skilSql,s.LRoleId,s.LSkillId,s.StrSkillCode,s.StrSkillName,s.NSkillLevel,s.NSkillValue,s.NSkillType,s.NAttackType,s.StrDesc,s.NCastTime,s.NDuration,s.StrEffectDesc)
				}
				if s.BChange && s.ChangeType == 2 {
					skilSql:="UPDATE tb_role_skill SET "
					skilSql+="nSkillLevel=?,"
					skilSql+="nSkillValue=?,"
					skilSql+="nCastTime=?,"
					skilSql+="nDuration=?"
					skilSql+="WHERE lId=?"
					tx.MustExec(skilSql,s.NSkillLevel,s.NSkillValue,s.NCastTime,s.NDuration,s.LId)
				}
			}
		}
	}
	tx.Commit()
}
