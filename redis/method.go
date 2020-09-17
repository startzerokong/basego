package redis

import (
	"time"
)

func Get(key string) (string, error) {
	rdb := GetRedisConnect()
	return rdb.Get(key).Result()
}

func Set(key, value string, expire time.Duration) (string, error) {
	rdb := GetRedisConnect()
	return rdb.Set(key, value, expire).Result()
}

func SetNx(key, value string, expire time.Duration) (bool, error) {
	rdb := GetRedisConnect()
	return rdb.SetNX(key, value, expire).Result()
}

func Del(key string)  {
	rdb := GetRedisConnect()
	rdb.Del(key)
}

func IncrBy(key string, value int64) (int64, error) {
	rdb := GetRedisConnect()
	return rdb.IncrBy(key, value).Result()
}

func Expire(key string, expire time.Duration)  {
	rdb := GetRedisConnect()
	rdb.Expire(key, expire)
}

func Exists(key string) (int64, error) {
	rdb := GetRedisConnect()
	return rdb.Exists(key).Result()
}

func ExpireAt(key string, expireAt time.Time)  {
	rdb := GetRedisConnect()
	rdb.ExpireAt(key, expireAt)
}

func HSet(key, hKey, value string) {
	rdb := GetRedisConnect()
	rdb.HSet(key, hKey, value)
}

func HGet(key, hKey string) (string, error) {
	rdb := GetRedisConnect()
	return rdb.HGet(key, hKey).Result()
}

func HMSet(key string, fields map[string]interface{}) (string, error) {
	rdb := GetRedisConnect()
	return rdb.HMSet(key, fields).Result()
}

//func HMGet(key string, fields ...string) {
//	rdb := GetRedisConnect()
//	rdb.HMGet(key, fields)
//}

func HGetAll(key string) (map[string]string, error) {
	rdb := GetRedisConnect()
	return rdb.HGetAll(key).Result()
}
