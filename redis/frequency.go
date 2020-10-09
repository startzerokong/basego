package redis

import (
	"github.com/startzerokong/basego/util"
	"time"
)

const FrequencyKey = "FREQUENCY::"

func IncrByIp(ip, fileName string) int64 {
	key := FrequencyKey + fileName + ip
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
