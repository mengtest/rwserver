package dao

import (
	"../../rbstruct/config"
	"../../rbwork/base"
	"../../rbwork/db"
	)

func GetLevelConfig() []config.LevelConfig {
	levels:=[]config.LevelConfig{}
	sqlc:="SELECT * FROM tb_level_config"
	base.LogInfo("SQL:",sqlc," Param:")
	err:=db.DB.Select(&levels,sqlc)
	base.CheckErr(err)
	return levels
}