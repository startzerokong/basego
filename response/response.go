package response

import (
	"github.com/gin-gonic/gin"
	"github.com/startzerokong/basego/define"
	"net/http"
)

func SuccessResponse(ctx *gin.Context, data interface{})  {
	ctx.Set("response_body", define.Response{
		Code:    100000,
		Message: "success",
		Data:    data,
	})
	Response(ctx)
}

func WrongResponse(ctx *gin.Context, errorCode int, errorMsg string)  {
	ctx.Set("response_body", define.Response{
		Code: errorCode,
		Message: errorMsg,
		Data: map[string]interface{}{},
	})
	Response(ctx)
}

func Response(ctx *gin.Context)  {
	if !ctx.IsAborted() {
		result,_ := ctx.Get("response_body")
		ctx.AbortWithStatusJSON(http.StatusOK, result)
	}
}
