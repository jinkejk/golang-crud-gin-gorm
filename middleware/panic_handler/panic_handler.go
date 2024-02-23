package panic_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang-crud-gin/data/response"
	"net/http"
)

func GlobalPanicHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic %v \n", r)
				webResponse := response.Response{
					Code:   1,
					Status: "Error",
					Msg:    fmt.Sprintf("panic %v \n", r),
					Data:   nil,
				}
				ctx.Header("Content-Type", "application/json")
				ctx.JSON(http.StatusOK, webResponse)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
