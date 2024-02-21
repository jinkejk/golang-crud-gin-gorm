package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler func(ctx *gin.Context) (data interface{}, err error)

// 中间件：处理异常和封装响应结果，同时适配gin.HandlerFunc
func ResponseHandler(h handler) gin.HandlerFunc {
	type response struct {
		Code    int
		Message string
		Data    interface{}
	}
	return func(ctx *gin.Context) {
		data, err := h(ctx)
		if err != nil {
			ctx.Error(err)
			return
		}
		resp := response{
			Code:    2000000,
			Message: "success",
			Data:    data,
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
