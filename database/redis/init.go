package redis

import (
	"crawler/lib"
	"github.com/gomodule/redigo/redis"
	"time"
)
var RedisClient *redis.Pool


func InitRedis() {
	var (
		host string
		auth string
		db   int
	)

	host = lib.Config.GetString("redis.host")
	auth = lib.Config.GetString("redis.auth")
	db = lib.Config.GetInt("redis.db")

	RedisClient = &redis.Pool{
		MaxIdle:     lib.Config.GetInt("redis.conn_max_idle"),
		MaxActive:   lib.Config.GetInt("redis.conn_max_active"),
		IdleTimeout: lib.Config.GetDuration("redis.conn_idle_timeout") * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host, redis.DialPassword(auth), redis.DialDatabase(db))
			if nil != err {
				panic("redis connect error:"+err.Error())
				return nil, err
			}
			lib.Log.Info("redis connect has ok")
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			if nil != err {
				panic("redis ping error:"+err.Error())
			}
			return err
		},
	}
}