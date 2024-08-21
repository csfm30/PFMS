package database

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/patcharp/golib/cache"
)

var caching cache.Redis

func InitCaching() error {
	caching = cache.NewWithCfg(cache.Config{
		Host:     viper.GetString("cache.host"),
		Port:     viper.GetString("cache.port"),
		Password: "",
		Db:       viper.GetInt("cache.db"),
	})

	fmt.Println("Connecting to redis cache", viper.GetString("cache.host"))
	if err := caching.Ping(); err != nil {
		return err
	}
	return nil
}

func CachingCtx() *cache.Redis {
	return &caching
}
