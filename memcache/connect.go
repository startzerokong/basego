package memcache

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/startzerokong/basego/util"
)

func InitMemCache() *memcache.Client {
	return GetMemCacheConnect()
}

func GetMemCacheConnect() *memcache.Client {
	config := util.GetMemCacheConfig()
	return memcache.New(config.Host + ":" + config.Port)
}