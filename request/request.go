package request

import (
	"github.com/gin-gonic/gin"
	"github.com/startzerokong/basego/response"
)

func GetQuery(ctx *gin.Context, name string) string {
	ret := ctx.Query(name)

	if len(ret) == 0 {
		response.WrongResponse(ctx, -1, "param error")
	}
	return ret
}

func GetQueryWithDefault(ctx *gin.Context, name, defaultParam string) string {
	if len(defaultParam) == 0 {
		response.WrongResponse(ctx, -1, "default param is null")
	}

	return ctx.DefaultQuery(name, defaultParam)
}

func GetPost(ctx *gin.Context, name string) string {
	ret := ctx.PostForm(name)

	if len(ret) == 0 {
		response.WrongResponse(ctx, -1, "param error")
	}

	return ret
}

func GetForm(ctx *gin.Context, name string) string {
	ret := ctx.PostForm(name)

	if len(ret) == 0 {
		response.WrongResponse(ctx, -1, "param error")
	}

	return ret
}

func GetFormWithDefault(ctx *gin.Context, name, defaultParam string) string {
	if len(defaultParam) == 0 {
		response.WrongResponse(ctx, -1, "default param is null")
	}

	return ctx.DefaultPostForm(name, defaultParam)
}
