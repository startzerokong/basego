package memcache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

func InitMemCache()  {
	memcache.New()
}