package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/startzerokong/basego/util"
	"time"
)

func GetRedisConnect() *redis.Client {
	redisConfig := util.GetRedisConfig()

	return redis.NewClient(&redis.Options{
		DB:           redisConfig.Db,
		Addr:         fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		DialTimeout:  time.Duration(redisConfig.ConnectTimeout*1000) * time.Millisecond,
		ReadTimeout:  time.Duration(redisConfig.ReadTimeout*1000) * time.Millisecond,
		WriteTimeout: time.Duration(redisConfig.WriteTimeout*1000) * time.Millisecond,
		Password:     redisConfig.PassWord,
	})
}
