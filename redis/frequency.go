package redis

import (
	"github.com/startzerokong/basego/util"
	"os"
	"time"
)

const FrequencyKey = "FREQUENCY::"

func IncrByIp(ip string) int64 {
	filename ,_ := os.Getwd()
	key := FrequencyKey + filename + ip
	ret, err := IncrBy(key, 1)
	if err != nil {
		return 0
	}
	if ret == 1 {
		config := util.GetIpFrequencyConfig()
		expire := time.Minute * time.Duration(config.Expire)
		Expire(key, expire)
	}
	return ret
}
