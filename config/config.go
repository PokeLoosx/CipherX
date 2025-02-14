package config

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"CipherX/config/autoload"
)

type Configuration struct {
	DB    autoload.DB    `mapstructure:"db" json:"db" yaml:"db"`
	Redis autoload.Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}

var (
	GinConfig Configuration
	GinDB     *gorm.DB
	GinRedis  *redis.Client
	GinLOG    *zap.Logger
	GinVP     *viper.Viper
)
