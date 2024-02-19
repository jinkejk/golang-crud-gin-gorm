package model

import "time"

type Tags struct {
	Id   int    `gorm:"column:id;type:int;primaryKey"`
	Name string `gorm:"column:name;type:varchar(255)"`
}

type Requests struct {
	Id            int       `gorm:"column:id;type:int;autoIncrement;primaryKey"`
	Workflow      string    `gorm:"column:workflow;type:longtext"`
	CurrentNode   string    `gorm:"column:current_node;type:varchar(25)"`
	Status        string    `gorm:"column:status;type:varchar(25);not null"`
	CreateTime    time.Time `gorm:"column:createTime;type:datetime;"`
	FinishTime    time.Time `gorm:"column:finishTime;type:datetime;"`
	RequestTypeId int       `gorm:"column:request_type_id;type:int;"`
	Params        string    `gorm:"column:params;type:longtext"`
	ProcessorId   int       `gorm:"column:processor_id;type:int;"`
	RequesterId   int       `gorm:"column:requester_id;type:int;"`
	WorkflowEnd   byte      `gorm:"column:workflowEnd;type:tinyint;not null"`
	BpmnId        int       `gorm:"column:bpmn_id;type:int;"`
}
