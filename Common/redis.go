package Common

import (
	"github.com/go-redis/redis"
	"fmt"
)
const (
	RedisKeyUserAuth string = "UserAuth:"
)

type initRedis struct {
	Client *redis.Client
}

func NewRedis() *initRedis {
	sysConfig := Microservice{}
	sysConfig.LoadConfig() //加载当前服务的配置
	addr := fmt.Sprintf("%s:%d", sysConfig.Cache.Host, sysConfig.Cache.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: sysConfig.Cache.Pwd,
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic("连接Redis失败:" + err.Error())
	}
	return &initRedis{
		Client: client,
	}
}

