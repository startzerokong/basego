package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/startzerokong/basego/encryption"
	"time"
)

func InitRequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := RequestId()

		ctx.Set("request_id", requestId)

		ctx.Next()
	}
}

func GetRequestId(ctx *gin.Context) string {
	requestId, exists := ctx.Get("request_id")
	if !exists {
		return ""
	}

	return requestId.(string)
}

func RequestId() string {
	chars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()+-")

	randomByte := []byte{
		chars[RangeRand(0, 73)], chars[RangeRand(0, 73)], chars[RangeRand(0, 73)], chars[RangeRand(0, 73)], chars[RangeRand(0, 73)],
	}

	randomData := fmt.Sprintf("%s-%s", Uniqid(""), string(randomByte))

	result := encryption.DoMd5(randomData)

	return result
}

func Uniqid(prefix string) string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", prefix, now.Unix(), now.UnixNano()%0x100000)
}
