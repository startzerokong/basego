package frequency

import (
	"github.com/gin-gonic/gin"
	"github.com/startzerokong/basego/define"
	"github.com/startzerokong/basego/redis"
	"github.com/startzerokong/basego/response"
	"github.com/startzerokong/basego/util"
)

func CheckFrequency() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		realIp := util.RealIp(ctx)
		fileName := ctx.Request.URL.Path
		count := redis.IncrByIp(realIp, fileName)
		config := util.GetIpFrequencyConfig()
		limit := config.Limit
		if count > limit {
			response.WrongResponse(ctx, 9999999, define.Message("9999999"))
		}
	}
}