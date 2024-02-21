package router

import (
	"github.com/gin-gonic/gin"
	"golang-crud-gin/controller"
)

func requestNewRouter(baseRouter *gin.RouterGroup) {
	var requestController = controller.NewRequestModelController()

	tagsRouter := baseRouter.Group("/requests")
	{
		tagsRouter.GET("", requestController.FindAll)
		tagsRouter.GET("/:requestId", requestController.FindById)
		tagsRouter.POST("", requestController.Create)
		tagsRouter.PATCH("/:requestId", requestController.Update)
		tagsRouter.DELETE("/:requestId", requestController.Delete)
	}
}
