package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var prefix = "{crawler_}"

func makeKey(key string) string{
	return fmt.Sprintf("%s_%s", prefix, key)
}

func  SetKV(key string, val interface{}) (int, error){
	redisClient := RedisClient.Get()
	defer redisClient.Close()
	key = makeKey(key)
	return redis.Int(redisClient.Do("set", key, val))
}

func IsExist(key string) (bool,error) {
	redisClient := RedisClient.Get()
	defer redisClient.Close()
	key = makeKey(key)
	return redis.Bool(redisClient.Do("get" ,key))
}
