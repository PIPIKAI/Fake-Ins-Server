package core

import (
	"os"

	"github.com/spf13/viper"
)

func InitViper() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("read config file err: " + err.Error())
	}

}
