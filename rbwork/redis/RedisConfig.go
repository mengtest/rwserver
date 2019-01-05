package redis

import (
	"gopkg.in/redis.v6"
)


var Client *redis.Client

func InitRedis(host string,password string){
	Client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
		PoolSize: 10,
	})
}

