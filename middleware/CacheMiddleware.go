package middleware

import (
	"github.com/chenyahui/gin-cache/persist"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// 单例模式
var cacheStore *persist.RedisStore

func buildCacheStore() {
	cacheStore = persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network: viper.GetString("redisconfig.network"),
		Addr:    viper.GetString("redisconfig.host"),
	}))
}

func GetCacheStore() *persist.RedisStore {
	if cacheStore == nil {
		buildCacheStore()
	}
	return cacheStore
}
