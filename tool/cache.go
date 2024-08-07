package tool

import (
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

type Cache struct {
}

var CacheObj = Cache{}

func (c *Cache) Init(redisCon redis.Options) {
	RedisClient = redis.NewClient(&redisCon)
}
