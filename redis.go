package util

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var redispool *redis.Pool

func redisMain() {
	InitRedis()
	key := "z"
	RedisDoZincrby(key, 3, "a")
	r, _ := RedisDoZrange(key, 0, -1, true)
	Ln(r)
}

func RedisDoZincrby(key string, increment int64, member interface{}) (err error) {
	conn := redispool.Get()
	defer conn.Close()
	_, err = conn.Do("ZINCRBY", key, increment, member)
	return
}

func RedisDoZrange(key string, start int, stop int, withscores bool) (result interface{}, err error) {
	conn := redispool.Get()
	defer conn.Close()
	if withscores {
		result, err = conn.Do("ZRANGE", key, start, stop, "withscores")
	} else {
		result, err = conn.Do("ZRANGE", key, start, stop)
	}
	return
}

func RedisDoGET(key string) (result interface{}, err error) {
	conn := redispool.Get()
	defer conn.Close()
	result, err = conn.Do("GET", key)
	return
}

func newRedisPool(server string, maxidle int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxidle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func InitRedis() {
	rdshost := "localhost:6371"
	rdsmaxpool := 100
	redispool = newRedisPool(rdshost, rdsmaxpool)
}
