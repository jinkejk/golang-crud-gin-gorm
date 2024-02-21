package controller

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func NewRequestModelController() *RequestModelController {
	return &RequestModelController{
		requestCurd: service.CreateRequestCurdFactory(),
	}
}

type RequestModelController struct {
	requestCurd *service.RequestsCurd
}

// Create Tags		godoc
// @Summary			Create tags
// @Description		Save tags data in Db.
// @Param			tags body request.CreateTagsRequest true "Create tags"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [post]
func (controller *RequestModelController) Create(ctx *gin.Context) {
	log.Info().Msg("create request")
	createTagsRequest := request.CreateRequestsModelRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	err = controller.requestCurd.Save(&createTagsRequest)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Update Tags		godoc
// @Summary			Update tags
// @Description		Update tags data.
// @Param			tagId path string true "update tags by id"
// @Param			tags body request.CreateTagsRequest true  "Update tags"
// @Tags			tags
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagId} [patch]
func (controller *RequestModelController) Update(ctx *gin.Context) {
	log.Info().Msg("update request")
	updateTagsRequest := request.UpdateRequestsModelRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	requestId := ctx.Param("requestId")
	id, err := strconv.Atoi(requestId)
	helper.ErrorPanic(err)
	updateTagsRequest.Id = int64(id)

	err = controller.requestCurd.Update(&updateTagsRequest)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// Delete Tags		godoc
// @Summary			Delete tags
// @Description		Remove tags data by id.
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagID} [delete]
func (controller *RequestModelController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete request")
	requestId := ctx.Param("requestId")
	id, err := strconv.Atoi(requestId)
	helper.ErrorPanic(err)

	err = controller.requestCurd.Delete(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindById Tags 		godoc
// @Summary				Get Single tags by id.
// @Param				tagId path string true "update tags by id"
// @Description			Return the tahs whoes tagId valu mathes id.
// @Produce				application/json
// @Tags				tags
// @Success				200 {object} response.Response{}
// @Router				/tags/{tagId} [get]
func (controller *RequestModelController) FindById(ctx *gin.Context) {
	log.Info().Msg("find by id requests")
	requestId := ctx.Param("requestId")
	id, err := strconv.Atoi(requestId)
	helper.ErrorPanic(err)

	res, err := controller.requestCurd.FindById(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   res,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll Tags 		godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			tags
// @Success			200 {obejct} response.Response{}
// @Router			/tags [get]
func (controller *RequestModelController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll requests")
	res, err := controller.requestCurd.FindAll()
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   res,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
