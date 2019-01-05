package redis

import (
	"../base"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"time"
)

const (
	RedisURL            = "redis://*****:6379"
	redisMaxIdle        = 3 //最大空闲连接数
	redisIdleTimeoutSec = 3 //最大空闲连接时间
	RedisPassword       = "*****"
)

var pool *redis.Pool

// redis连接池
func InitRedisPool(redisURL string, password string) {
	base.LogError("redis connecting...")
	pool = &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
			if err != nil {
				base.LogInfo("redis connection error")
				return nil, err
			}
			//验证redis密码
			if _, authErr := c.Do("AUTH", password); authErr != nil {
				base.LogInfo("redis auth password error: ")
				return nil, authErr
			}
			base.LogInfo("redis connect success")
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				base.LogError("ping redis error")
				return err
			}
			return nil
		},
	}
}

func Set(k, v string) error {
	c := pool.Get()
	_, err := c.Do("SET", k, v)
	return err
}

func GetStringValue(k string) string {
	c := pool.Get()
	defer c.Close()
	value, err := redis.String(c.Do("GET", k))
	if err != nil {
		fmt.Println("Get Error: ", err.Error())
		return ""
	}
	return value
}

func SetKeyExpire(k string, ex int) {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", k, ex)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func CheckKey(k string) bool {
	c := pool.Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("EXISTS", k))
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return exist
	}
}

func DelKey(k string) error {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("DEL", k)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SetJson(k string, data interface{}) error {
	c := pool.Get()
	defer c.Close()
	value, _ := json.Marshal(data)
	n, _ := c.Do("SETNX", k, value)
	if n != int64(1) {
		return errors.New("set failed")
	}
	return nil
}

func GetByte(key string) ([]byte, error) {
	c := pool.Get()
	defer c.Close()
	value, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return value, nil
}

//---------Set集合---------------
func Sadd(k string, v ...string) error {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("SADD", k, v)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Smembers(k string) ([]string, error) {
	c := pool.Get()
	defer c.Close()
	value, err := redis.Strings(c.Do("SMEMBERS", k))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return value, nil
}

func Scard(k string) int {
	c := pool.Get()
	defer c.Close()
	value, err := redis.Int(c.Do("SCARD", k))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return value
}

//---------------------------------------------
