package util

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/garyburd/redigo/redis"
)

// 声明一些全局变量
var (
	redisPool   *redis.Pool
	redisServer = "127.0.0.1:32768"
)

func init() {
	redisPool = newPool(redisServer, "redispw")
}

// 初始化一个pool
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

//func LpushRedis(key string, expriedTime int, args ...interface{}) {
//	conn := redisPool.Get()
//	defer conn.Close()
//	conn.Do("lpush", key, args)
//	if expriedTime > 0 {
//		conn.Do("EXPIRE", key, expriedTime)
//	}
//
//}

func SetStringToRedis(key string, expriedTime int, args string) {
	conn := redisPool.Get()
	defer conn.Close()
	_, err := conn.Do("set", key, args)
	if err != nil {
		fmt.Println(err)
	}
	if expriedTime > 0 {
		conn.Do("EXPIRE", key, expriedTime)
	}
}

func DelKey(key string) {
	conn := redisPool.Get()
	defer conn.Close()
	conn.Do("del", key)
}

func GetStringFromRedis(key string) string {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	conn := redisPool.Get()
	defer conn.Close()

	//value, err := redis.String(conn.Do("GET", key))
	//content,err := conn.Do("get", key)
	res, err := conn.Do("GET", key)

	if res != nil {
		value, err := redis.String(res, err)

		if err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
		return value
	}
	return ""

}

func GetExpiredTimeFromRedis(key string) int {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	conn := redisPool.Get()
	defer conn.Close()

	value, err := redis.Int(conn.Do("TTL", key))
	//content,err := conn.Do("get", key)
	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
	}
	return value

}

//existKey, _ := redis.Bool(conn.Do("EXISTS", key))
//content,err := conn.Do("lindex", key, i)
//if err != nil {
//fmt.Println(" err:",err)
//}
//func AddTest(id int, name string) {
//	conn := redisPool.Get()
//	defer conn.Close()
//	//redis操作
//	v, err := conn.Do("SET", id, name)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(v)
//	v, err = redis.String(conn.Do("GET", id))
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(v)
//
//}

func GetRedisConn() redis.Conn {
	return redisPool.Get()
}

func Incr(key string) int {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	conn := redisPool.Get()
	defer conn.Close()

	value, err := redis.Int(conn.Do("Incr", key))
	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
	}
	return value

}
