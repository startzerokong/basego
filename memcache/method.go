package memcache

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/startzerokong/basego/util"
)

func Get(key string) string {
	client := GetMemCacheConnect()
	item, err := client.Get(key)
	if err != nil {
		return ""
	}
	return util.ByteToString(item.Value)
}

func Set(key, value string) bool {
	client := GetMemCacheConnect()
	err := client.Set(&memcache.Item{Key: key, Value: util.StringToByte(value)})
	if err != nil {
		return false
	}
	return true
}

func Del(key string) bool {
	client := GetMemCacheConnect()
	err := client.Delete(key)
	if err != nil {
		return false
	}
	return true
}

func Add(key, value string) bool {
	client := GetMemCacheConnect()
	err := client.Add(&memcache.Item{Key: key, Value: util.StringToByte(value)})
	if err != nil {
		return false
	}
	return true
}

func Incr(key string, value uint64) (uint64, error) {
	client := GetMemCacheConnect()
	ret, err := client.Increment(key, value)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func Decr(key string, value uint64) (uint64, error) {
	client := GetMemCacheConnect()
	ret, err := client.Decrement(key, value)
	if err != nil {
		return 0, err
	}
	return ret, nil
}
