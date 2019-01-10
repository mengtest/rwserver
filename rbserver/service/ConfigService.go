package service

import (
	"../dao"
	"../util"
)

//加载等级配置
func InitLevel()  {
	levels:=dao.GetLevelConfig()
	for _,level:= range levels  {
		util.LevelMap[level.NLevel]=&level
	}
}