package service

import (
	"../../rbwork/base"
	"../../rbwork/db"
	"database/sql"
	"time"
)

//定义用户结构体
type User struct {
    LId int64 `db:"lId" json:"lId"`
    StrName string `db:"strName" json:"strName"`
    StrPwd string `db:"strPwd" json:"-"`  //json格式化的时候不返回
	StrRealName string `db:"strRealName" json:"strRealName"`
	StrIdCardNo string `db:"strIdCardNo" json:"strIdCardNo"`
	NAuthStatus int `db:"nAuthStatus" json:"nAuthStatus"`
	StrMobile string `db:"strMobile" json:"strMobile"`
	StrEmail  string `db:"strEmail" json:"strEmail"`
	DtUpdateTime string `db:"dtUpdateTime" json:"dtUpdateTime"`
	DtCreateTime string `db:"dtCreateTime" json:"dtCreateTime"`
	NDeleted int `db:"nDeleted" json:"-"`
}

type UserData struct {
	Token string `json:"token"`
	User User `json:"user"`
}

func GetUserByMobile(mobile string) (User,int) {
	user:=User{}
	sqlc:="SELECT * FROM tb_user WHERE nDeleted=0 AND strMobile=? ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sqlc," Param:",mobile)
	err:=db.DB.Get(&user,sqlc,mobile)
	if err==sql.ErrNoRows{
		return user,0
	}else if base.CheckErr(err) {
		return user,-1
	}
   return user,1
}

func GetUserById(lId int64) (User,int) {
	user:=User{}
	sqlc:="SELECT * FROM tb_user WHERE nDeleted=0 AND lId=? ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sqlc," Param:",lId)
	err:=db.DB.Get(&user,sqlc,lId)
	if err==sql.ErrNoRows{
		return user,0
	}else if base.CheckErr(err) {
		return user,-1
	}
	return user,1
}

func GetUserByName(name string) (User,int) {
	user:=User{}
	sqlc:="SELECT * FROM tb_user WHERE nDeleted=0 AND strName=? ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sqlc," Param:",name)
	err:=db.DB.Get(&user,sqlc,name)
	if err==sql.ErrNoRows{
		return user,0
	}else if base.CheckErr(err) {
		return user,-1
	}
	return user,1
}

func SaveUser(user User) int {
	sqlc:="INSERT INTO tb_user (strName,strPwd,dtUpdateTime,dtCreateTime) VALUES (?,?,?,?)"
	res,err:=db.DB.Exec(sqlc,user.StrName,user.StrPwd,time.Now(),time.Now())
	if base.CheckErr(err) {
		return -1
	}
	ret,_:=res.RowsAffected()
	return int(ret)
}
