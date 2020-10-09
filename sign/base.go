package sign

import (
	"github.com/gin-gonic/gin"
	"github.com/startzerokong/basego/define"
	"github.com/startzerokong/basego/encryption"
	"github.com/startzerokong/basego/request"
	"github.com/startzerokong/basego/response"
)

const salt = "let-eat-together"

func BuildUrl(ctx *gin.Context) string {
	queryString := ctx.Request.URL.Query()

	var query string
	for index, value := range queryString {
		for _,value1 := range value{
			if len(value1) != 0 {
				query += index
				query += "&"
				query += value1
			}
		}
	}

	return query
}

func BuildUrlWithSalt(ctx *gin.Context) string {
	return encryption.DoMd5(BuildUrl(ctx) + salt)
}

func CheckSign() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sign := request.GetQuery(ctx, "sign")

		if len(sign) == 0 {
			response.WrongResponse(ctx, 9999998, define.Message("9999998"))
		}

		buildSign := BuildUrlWithSalt(ctx)

		if sign != buildSign {
			response.WrongResponse(ctx, 9999998, define.Message("9999998"))
		}
	}
}
