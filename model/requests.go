package model

import (
	"fmt"
	"time"
)

func CreateRequestsModelFactory() *RequestsModel {
	return &RequestsModel{BaseModel: BaseModel{DB: DatabaseConnection()}}
}

type RequestsModel struct {
	BaseModel
	Workflow      string    `gorm:"column:workflow;type:longtext"`
	CurrentNode   string    `gorm:"column:current_node;type:varchar(25)"`
	Status        string    `gorm:"column:status;type:varchar(25);not null"`
	CreateTime    time.Time `gorm:"column:createTime;type:datetime;"`
	FinishTime    time.Time `gorm:"column:finishTime;type:datetime;"`
	RequestTypeId int       `gorm:"column:request_type_id;type:int;"`
	Params        string    `gorm:"column:params;type:longtext"`
	ProcessorId   *int      `gorm:"column:processor_id;type:int;"`
	RequesterId   *int      `gorm:"column:requester_id;type:int;"`
	WorkflowEnd   byte      `gorm:"column:workflowEnd;type:tinyint;not null"`
	BpmnId        int       `gorm:"column:bpmn_id;type:int;"`
}

func (t *RequestsModel) TableName() string {
	return "request_request"
}

// Delete 按照ID删除
func (t *RequestsModel) Delete(requestId int) error {
	sql := "DELETE FROM  request_request WHERE id=? "
	err := t.Exec(sql, requestId).Error
	return err
}

// FindAll 查找
func (t *RequestsModel) FindAll() ([]*RequestsModel, error) {
	var objects []*RequestsModel
	sql := `select id,current_node,status,createTime,finishTime,request_type_id,params,processor_id,requester_id,workflowEnd,bpmn_id 
		    from request_request limit 100`
	err := t.Raw(sql).Find(&objects).Error
	if err == nil {
		return objects, nil
	} else {
		return nil, err
	}
}

// FindById 查找
func (t *RequestsModel) FindById(requestId int) (*RequestsModel, error) {
	sql := `select id,current_node,status,createTime,finishTime,request_type_id,params,processor_id,requester_id,workflowEnd,bpmn_id 
		    from request_request where id=? limit 1`
	result := t.Raw(sql, requestId).First(t)
	if result.Error == nil {
		// 账号密码验证成功
		return t, nil
	} else {
		fmt.Printf("根据账号查询单条记录出错: %v", result.Error)
	}
	return nil, nil
}

// Save 创建
func (t *RequestsModel) Save(object *RequestsModel) error {
	err := t.Create(object).Error
	return err
	//sql := "INSERT INTO request_request(current_node,status,createTime,finishTime,request_type_id,params,processor_id,requester_id,workflowEnd,bpmn_id) values (?,?,?,?,?,?,?,?,?,?)"
	//if t.Exec(sql, object.CurrentNode, object.Status, object.CreateTime, object.FinishTime, object.RequestTypeId, object.Params,
	//	object.ProcessorId, object.RequesterId, object.WorkflowEnd, object.BpmnId).RowsAffected > 0 {
	//	return true
	//}
	//return false
}

// Update 更新
func (t *RequestsModel) Update(object *RequestsModel) error {
	err := t.Updates(object).Error
	return err
}
