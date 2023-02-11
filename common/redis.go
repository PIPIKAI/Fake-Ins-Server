package common

import (
	"log"

	"github.com/gin-contrib/sessions/redis"
	"github.com/spf13/viper"
)

var store redis.Store

func InitRedis() redis.Store {
	host := viper.GetString("redisconfig.host")
	network := viper.GetString("redisconfig.network")
	conneters := viper.GetInt("redisconfig.conneters")
	encodePassword := viper.GetString("redisconfig.network")
	password := viper.GetString("redisconfig.password")

	thestore, err := redis.NewStore(conneters, network, host, password, []byte(encodePassword))
	if err != nil {
		panic("redis 连接失败")
	} else {
		log.Println("redis 连接成功")
	}
	store = thestore
	return store
}
func GetRedis() redis.Store {
	return store
}
