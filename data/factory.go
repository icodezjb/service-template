package data

import "github.com/buchenglei/service-template/data/redis"

func NewRedisData() RedisData {
	return redis.New()
}
