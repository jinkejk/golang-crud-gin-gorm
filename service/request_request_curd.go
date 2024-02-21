package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/model"
)

func CreateRequestCurdFactory() *RequestsCurd {
	return &RequestsCurd{model.CreateRequestsModelFactory()}
}

type RequestsCurd struct {
	requestsModel *model.RequestsModel
}

// Delete 按照ID删除
func (t *RequestsCurd) Delete(requestId int) error {
	return t.requestsModel.Delete(requestId)
}

// FindAll 查找
func (t *RequestsCurd) FindAll() ([]response.RequestsModelResponse, error) {
	res, err := t.requestsModel.FindAll()
	if err == nil {
		var objs []response.RequestsModelResponse
		for _, value := range res {
			rmr := response.RequestsModelResponse{}
			response.RequestConvertToResponse(&rmr, value)
			objs = append(objs, rmr)
		}
		return objs, nil
	} else {
		return nil, err
	}

}

// FindById 查找
func (t *RequestsCurd) FindById(requestId int) (*response.RequestsModelResponse, error) {
	obj, err := t.requestsModel.FindById(requestId)
	if err == nil {
		rmr := response.RequestsModelResponse{}
		response.RequestConvertToResponse(&rmr, obj)
		return &rmr, nil
	} else {
		return nil, err
	}
}

// Save 创建
func (t *RequestsCurd) Save(object *request.CreateRequestsModelRequest) error {
	rm := model.RequestsModel{}
	request.CreateRequestConvertToModel(object, &rm)
	return t.requestsModel.Save(&rm)
}

// Update 更新
func (t *RequestsCurd) Update(object *request.UpdateRequestsModelRequest) error {
	rm := model.RequestsModel{}
	request.UpdateRequestConvertToModel(object, &rm)
	return t.requestsModel.Update(&rm)
}
