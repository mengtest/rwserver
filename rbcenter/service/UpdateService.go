package service

import (
	"../../rbwork/base"
	"../../rbwork/db"
	"database/sql"
)

type Version struct {
	LId int64 `db:"lId" json:"lId"`
	StrName string `db:"strName" json:"strName"`
	StrVersion  string `db:"strVersion" json:"strVersion"`
	StrPath     string `db:"strPath" json:"strPath"`
	StrMd5 string `db:"strMd5" json:"strMd5"`
	NAppType  int `db:"nAppType" json:"nAppType"`
	DtUpdateTime sql.NullString `db:"dtUpdateTime" json:"-"`
	DtCreateTime sql.NullString `db:"dtCreateTime" json:"-"`
	NDeleted int `db:"nDeleted" json:"-"`
}

func CheckVersion() []Version {
	version := []Version{}
	sql:="SELECT * FROM tb_version WHERE nDeleted=0 GROUP BY nAppType ORDER BY dtCreateTime DESC"
	base.LogInfo("SQL:",sql)
	err:=db.DB.Select(&version,sql)
	base.CheckErr(err)
	return version
}
