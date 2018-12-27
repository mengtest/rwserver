package service

import (
	"../../rbwork/base"
	"../../rbwork/db"
)
//定义版本结构体
type Version struct {
	LId int64 `db:"lId" json:"lId"`
	StrName string `db:"strName" json:"strName"`
	StrVersion  string `db:"strVersion" json:"strVersion"`
	StrPath     string `db:"strPath" json:"strPath"`
	StrMd5 string `db:"strMd5" json:"strMd5"`
	NAppType  int `db:"nAppType" json:"nAppType"`
	DtUpdateTime string `db:"dtUpdateTime" json:"-"`
	DtCreateTime string `db:"dtCreateTime" json:"-"`
	NDeleted int `db:"nDeleted" json:"-"`
}

//对sql.NullString进行处理
func GetNewVersion(v Version) Version {
	return v
}

func CheckVersion() []Version {
	version := []Version{}
	sql:="SELECT * FROM tb_version WHERE nDeleted=0 GROUP BY nAppType ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sql)
	err:=db.DB.Select(&version,sql)
	base.CheckErr(err)
	return version
}
