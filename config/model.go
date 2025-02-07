package config

import (
	"go.uber.org/zap"
)

type CFG struct {
	Apikey string `mapstructure:"APIKEY" yaml:"APIKEY"`
	Domain string `mapstructure:"DOMAIN" yaml:"DOMAIN"`
	Log    *zap.Logger
	DB     DB `yaml:"DB" yaml:"DB"`
}

type DB struct {
	Host  string `mapstructure:"HOST" yaml:"HOST"`
	Name  string `mapstructure:"NAME" yaml:"NAME"`
	Login string `mapstructure:"LOGIN" yaml:"LOGIN"`
	Pass  string `mapstructure:"PASS" yaml:"PASS"`
	Table string `mapstructure:"TABLE" yaml:"TABLE"`
}
