package util

import "github.com/gin-gonic/gin"

func RealIp(ctx *gin.Context) string {
	return ctx.ClientIP()
}
