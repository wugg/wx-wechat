package tool

import (
	"github.com/go-redis/redis/v8"
	config2 "wechat/config"
)

var RedisClient = &redis.Client{}

func init() {
	RedisClient = redis.NewClient(&config2.RedisConf)
}
