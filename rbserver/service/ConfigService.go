package service

import (
	"../dao"
	Gloal "../util"
)

//加载等级配置
func InitLevel()  {
	levels:=dao.GetLevelConfig()
	for _,level:= range levels  {
		Gloal.LevelConfig[level.NLevel]=&level
	}
}