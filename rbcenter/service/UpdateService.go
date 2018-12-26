package service

import (
	"../../rbwork/base"
	"../../rbwork/db"
	"database/sql"
)

type Version struct {
	Id int64 `db:"id" json:"id"`
	PackageName string `db:"packageName" json:"packageName"`
	Version  string `db:"version" json:"version"`
	Path     string `db:"path" json:"path"`
	Md5 string `db:"md5" json:"md5"`
	AppType  int `db:"appType" json:"appType"`
	UpdateTime sql.NullString `db:"updateTime" json:"-"`
	CreateTime sql.NullString `db:"createTime" json:"-"`
	Delete int `db:"delete" json:"-"`
}

func CheckVersion() []Version {
	version := []Version{}
	sql:="SELECT * FROM tb_version WHERE delete=0 GROUP BY nType ORDER BY createTime DESC"
	base.LogInfo("SQL:",sql)
	db.DB.Select(&version,sql)
	return version
}
