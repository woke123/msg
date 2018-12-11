package main

import (
	"github.com/gomodule/redigo/redis"
	"msg/common"
	"msg/config"
)

func main(){
    config.InitMsg()
	redis_c, err := redis.Dial("tcp", config.C.Redis.Host, redis.DialPassword(config.C.Redis.Auth))
	if err != nil {
		return
	}
    common.InitEmail(redis_c)
}