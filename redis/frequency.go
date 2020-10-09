package redis

import "github.com/startzerokong/basego/util"

const FrequencyKey = "FREQUENCY::"

func IncrByIp(ip string) int64 {
	key := FrequencyKey + ip
	ret, err := IncrBy(key, 1)
	if err != nil {
		return 0
	}
	if ret == 1 {
		config := util.GetIpFrequencyConfig()
		expire := config.Expire
		Expire(key, expire)
	}
	return ret
}
