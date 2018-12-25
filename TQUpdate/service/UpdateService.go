package service

import (
	"../../TQBase/db"
	"database/sql"
)

type Version struct {
	LId int64 `db:"lId" json:"lId"`
	StrName string `db:"strName" json:"strName"`
	StrVersion  string `db:"strVersion" json:"strVersion"`
	StrPath     string `db:"strPath" json:"strPath"`
	StrMd5 string `db:"strMd5" json:"strMd5"`
	NType  int `db:"nType" json:"nType"`
	NDelete int `db:"nDelete" json:"-"`
	DtUpdateTime sql.NullString `db:"dtUpdateTime" json:"-"`
	DtCreateTime sql.NullString `db:"dtCreateTime" json:"-"`
}

func CheckVersion() []Version {
	version := []Version{}
	db.DB.Select(&version,"SELECT * FROM tb_version GROUP BY nType ORDER BY dtCreateTime DESC")
	return version
}
