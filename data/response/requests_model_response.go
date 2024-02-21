package response

import (
	"golang-crud-gin/model"
	"time"
)

type RequestsModelResponse struct {
	Id            int64     `json:"id"`
	CurrentNode   string    `json:"current_node"`
	Status        string    `json:"status"`
	CreateTime    time.Time `json:"create_time"`
	FinishTime    time.Time `json:"finish_time"`
	RequestTypeId int       `json:"request_type_id"`
	Params        string    `json:"params"`
	ProcessorId   *int      `json:"processor_id"`
	RequesterId   *int      `json:"requester_id"`
	WorkflowEnd   byte      `json:"workflow_end"`
	BpmnId        int       `json:"bpmn_id"`
}

func RequestConvertToResponse(r *RequestsModelResponse, rm *model.RequestsModel) {
	r.Id = rm.Id
	r.CurrentNode = rm.CurrentNode
	r.Status = rm.Status
	r.CreateTime = rm.CreateTime
	r.FinishTime = rm.FinishTime
	r.RequestTypeId = rm.RequestTypeId
	r.Params = rm.Params
	r.ProcessorId = rm.ProcessorId
	r.RequesterId = rm.RequesterId
	r.WorkflowEnd = rm.WorkflowEnd
	r.BpmnId = rm.BpmnId
}
