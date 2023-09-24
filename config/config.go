package config

type Server struct {

	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	Redis  Redis           `mapstructure:"redis" json:"redis" yaml:"redis"`

	Local Local `mapstructure:"local" json:"local" yaml:"local"`

	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
}
