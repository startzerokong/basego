package frequency

import (
	"github.com/gin-gonic/gin"
	"github.com/startzerokong/basego/redis"
	"github.com/startzerokong/basego/response"
	"github.com/startzerokong/basego/util"
)

func CheckSign() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		realIp := util.RealIp(ctx)
		count := redis.IncrByIp(realIp)
		config := util.GetIpFrequencyConfig()
		limit := config.Limit
		if count > limit {
			response.WrongResponse(ctx, 10000, "wrong")
		}
	}
}