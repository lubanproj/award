package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// GetConn returns a redis connection
func GetConn() (redis.Conn, error) {
	conn , err := redis.Dial(Conf.Redis.Network, fmt.Sprintf("%s:%d", Conf.Redis.Ip, Conf.Redis.Port))

	if err != nil {
		fmt.Println("connect to redis error ", err)
		return nil, err
	}

	return conn, nil
}

