package service

import (
	"../../rbwork/base"
	"../../rbwork/db"
	"database/sql"
)

//定义用户结构体
type User struct {
    LId int64 `db:"lId" json:"lId"`
    StrName string `db:"strName" json:"strName"`
    StrPwd string `db:"strPwd" json:"strPwd"`
	StrRealName string `db:"strRealName" json:"strRealName"`
	StrIdCardNo string `db:"strIdCardNo" json:"strIdCardNo"`
	NAuthStatus int `db:"nAuthStatus" json:"nAuthStatus"`
	StrMobile string `db:"strMobile" json:"strMobile"`
	StrEmail  string `db:"strEmail" json:"strEmail"`
	DtUpdateTime sql.NullString `db:"dtUpdateTime" json:"-"`
	DtCreateTime sql.NullString `db:"dtCreateTime" json:"-"`
	NDeleted int `db:"nDeleted" json:"-"`
}

func GetUserByMobile(mobile string) (User,error) {
	user:=User{}
	sql:="SELECT * FROM tb_user WHERE nDeleted=0 AND strMobile=? ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sql," Param:",mobile)
	err:=db.DB.Get(&user,sql,mobile)
	if base.CheckErr(err) {
		return user,err
	}
   return user,nil
}

func GetUserById(lId int64) (User,error) {
	user:=User{}
	sql:="SELECT * FROM tb_user WHERE nDeleted=0 AND lId=? ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sql," Param:",lId)
	err:=db.DB.Get(&user,sql,lId)
	if base.CheckErr(err) {
		return user,err
	}
	return user,nil
}

func GetUserByName(name string) (User,error) {
	user:=User{}
	sql:="SELECT * FROM tb_user WHERE nDeleted=0 AND strName=? ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sql," Param:",name)
	err:=db.DB.Get(&user,sql,name)
	if base.CheckErr(err) {
		return user,err
	}
	return user,nil
}


func Login(mobile string,pwd string) (User, int, string) {
	user:=User{}
	return user,0,""
}