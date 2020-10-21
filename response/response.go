package response

import (
	"github.com/gin-gonic/gin"
	"github.com/startzerokong/basego/define"
	"github.com/startzerokong/basego/util"
	"net/http"
)

func Response(ctx *gin.Context, data interface{}, errorCode int, errorMsg string, err error) {
	if err != nil {
		WrongResponse(ctx, errorCode, errorMsg)
	}
	SuccessResponse(ctx, data)
}

func SuccessResponse(ctx *gin.Context, data interface{})  {
	requestId := util.GetRequestId(ctx)
	ctx.Set("response_body", define.Response{
		RequestId:  requestId,
		Code:       100000,
		Message:    "success",
		Data:       data,
	})
	ResultResponse(ctx)
}

func WrongResponse(ctx *gin.Context, errorCode int, errorMsg string)  {
	requestId := util.GetRequestId(ctx)
	ctx.Set("response_body", define.Response{
		RequestId:  requestId,
		Code:       errorCode,
		Message:    errorMsg,
		Data:       map[string]interface{}{},
	})
	ResultResponse(ctx)
}

func ResultResponse(ctx *gin.Context)  {
	if !ctx.IsAborted() {
		result,_ := ctx.Get("response_body")
		ctx.AbortWithStatusJSON(http.StatusOK, result)
	}
}
