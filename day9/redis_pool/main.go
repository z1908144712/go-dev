package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379", redis.DialPassword("root"))
		},
	}
}

func main() {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
	pool.Close()
}
