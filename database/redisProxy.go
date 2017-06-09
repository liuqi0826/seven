package database

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type RedisProxy struct {
	Pool    *redis.Pool
	address string
	active  int
	idle    int
}

func (this *RedisProxy) RedisProxy(address string, active int, idle int) error {
	var err error
	this.address = address
	this.active = active
	this.idle = idle
	this.Pool = &redis.Pool{
		MaxIdle:   idle,
		MaxActive: active,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	fmt.Println("Redis Proxy Ready...")
	return err
}

func (this *RedisProxy) Ping() error {
	conn, err := redis.Dial("tcp", this.address)
	if err != nil {
		fmt.Println("RedisProxy Connect Error: ", err)
		return err
	}
	r, err := conn.Do("ping")
	if err != nil {
		fmt.Println("RedisProxy Ping Error: ", err)
		return err
	} else {
		if r == "PONG" {
			fmt.Println("Redis Proxy Ping Success.")
		}
	}
	return err
}

func (this *RedisProxy) GetConnect() redis.Conn {
	return this.Pool.Get()
}

func (this *RedisProxy) Close() error {
	var err error
	err = this.Pool.Close()
	if err != nil {
		fmt.Println("RedisProxy Close Error: ", err)
		return err
	}
	return err
}
