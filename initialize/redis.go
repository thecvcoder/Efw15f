package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/thecvcoder/cloud-api-go/global"
	"go.uber.org/zap"
)

func InitRedis() {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis连接失败:", zap.Error(err))
	} else {
		global.LOG.Info("redis连接成功:", zap.String("pong", pong))
		global.REDIS = client
	}
}
