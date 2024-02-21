package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang-crud-gin/controller"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"
	"golang-crud-gin/service"
)

func tagsNewRouter(baseRouter *gin.RouterGroup) {
	var (
		db             = model.DatabaseConnection()
		validate       = validator.New()
		tagsRepository = repository.NewTagsRepositoryImpl(db)
		tagsService    = service.NewTagsServiceImpl(tagsRepository, validate)
		tagsController = controller.NewTagsController(tagsService)
	)

	tagsRouter := baseRouter.Group("/tags")
	{
		tagsRouter.GET("", tagsController.FindAll)
		tagsRouter.GET("/:tagId", tagsController.FindById)
		tagsRouter.POST("", tagsController.Create)
		tagsRouter.PATCH("/:tagId", tagsController.Update)
		tagsRouter.DELETE("/:tagId", tagsController.Delete)
	}
}
