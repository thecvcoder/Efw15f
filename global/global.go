package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/thecvcoder/cloud-api-go/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG *config.Config
	LOG    *zap.SugaredLogger
	DB     *gorm.DB
	Viper  *viper.Viper
	REDIS  *redis.Client
)
