package router

import (
	"golang-crud-gin/middleware/panic_handler"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// global panic handler
	router.Use(panic_handler.GlobalPanicHandler())

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")

	// add sub router
	tagsNewRouter(baseRouter)
	requestNewRouter(baseRouter)

	return router
}
