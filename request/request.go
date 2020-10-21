package request

import (
	"github.com/gin-gonic/gin"
	"github.com/startzerokong/basego/define"
	"github.com/startzerokong/basego/response"
)

func GetQuery(ctx *gin.Context, name string) string {
	ret := ctx.Query(name)

	if len(ret) == 0 {
		response.WrongResponse(ctx, 9999997, define.Message("9999997"))
	}
	return ret
}

func GetQueryWithDefault(ctx *gin.Context, name, defaultParam string) string {
	if len(defaultParam) == 0 {
		response.WrongResponse(ctx, 9999996, define.Message("9999996"))
	}

	return ctx.DefaultQuery(name, defaultParam)
}

func GetPost(ctx *gin.Context, name string) string {
	ret := ctx.PostForm(name)

	if len(ret) == 0 {
		response.WrongResponse(ctx, 9999997, define.Message("9999997"))
	}

	return ret
}

func GetForm(ctx *gin.Context, name string) string {
	ret := ctx.PostForm(name)

	if len(ret) == 0 {
		response.WrongResponse(ctx, 9999997, define.Message("9999997"))
	}

	return ret
}

func GetFormWithDefault(ctx *gin.Context, name, defaultParam string) string {
	if len(defaultParam) == 0 {
		response.WrongResponse(ctx, 9999996, define.Message("9999996"))
	}

	return ctx.DefaultPostForm(name, defaultParam)
}

func GetHead(ctx *gin.Context, name string) string {
	ret := ctx.GetHeader(name)

	if len(ret) == 0 {
		response.WrongResponse(ctx, 9999997, define.Message("9999997"))
	}

	return ret
}
