package common

import (
	"log"
	"time"

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
	var thestore redis.Store
	var err error
	for {
		thestore, err = redis.NewStore(conneters, network, host, password, []byte(encodePassword))
		if err != nil {
			log.Println("[warn]: redis 连接失败 5s后尝试重新连接")
			time.Sleep(5 * time.Second)
		} else {
			log.Println("redis 连接成功")
			break
		}
	}

	store = thestore
	return store
}
func GetRedis() redis.Store {
	return store
}
