package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

// GetConn returns a redis connection
func GetConn() redis.Conn{
	conn , err := redis.Dial("tcp", fmt.Sprintf("%s:%d", "127.0.0.1", 6379))

	if err != nil {
		log.Println("connect to redis error ", err)
		return nil
	}

	return conn
}

