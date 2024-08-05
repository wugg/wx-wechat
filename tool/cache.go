package tool

import "github.com/go-redis/redis/v8"

var RedisClient *redis.Client
var redisConfig = &redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
}

func init() {
	RedisClient = redis.NewClient(redisConfig)
}
