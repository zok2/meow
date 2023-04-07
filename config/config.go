package config

type Server struct {

	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	System System          `mapstructure:"system" json:"system" yaml:"system"`

	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
}
