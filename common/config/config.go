package config

type Server struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Jwt    Jwt    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
