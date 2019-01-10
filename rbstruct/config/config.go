package config

//等级配置
type LevelConfig struct {
	LId          int64          `db:"lId" json:"lId"`
	NLevel       int            `db:"nLevel" json:"nLevel"`
	NExp         int64          `db:"nExp" json:"nExp"`
}
