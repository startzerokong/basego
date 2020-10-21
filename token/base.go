package token

import (
	"github.com/gin-gonic/gin"
	"github.com/startzerokong/basego/define"
	"github.com/startzerokong/basego/request"
	"github.com/startzerokong/basego/response"
	"github.com/startzerokong/basego/util"
	"strings"
	"time"
)

const (
	tokenExpireTime = 86400 * 30
	tokenSalt = "this is a default token salt"
)

func BuildUrl()  {

}

func CreateToken(uid string) string {
	timeObj := time.Now()
	unixTime := timeObj.Unix()
	expireTimeInt := unixTime + tokenExpireTime
	expireTime := util.Int64ToString(expireTimeInt)

	return uid + "." + expireTime
}

func ReverseToken(token string) []string {
	return strings.Split(token, ".")
}

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := request.GetHead(ctx, "token")
		uid := request.GetQuery(ctx, "uid")
		tokenSlice := ReverseToken(token)
		if tokenSlice[0] != uid {
			response.WrongResponse(ctx, 9999995, define.Message("9999995"))
		}
		nowTime := util.Int64ToString(time.Now().Unix())
		if nowTime > tokenSlice[1] {
			response.WrongResponse(ctx, 9999994, define.Message("9999994"))
		}

	}
}
