package request

import (
	"golang-crud-gin/model"
	"time"
)

type CreateRequestsModelRequest struct {
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

func CreateRequestConvertToModel(r *CreateRequestsModelRequest, rm *model.RequestsModel) {
	rm.CurrentNode = r.CurrentNode
	rm.Status = r.Status
	rm.CreateTime = r.CreateTime
	rm.FinishTime = r.FinishTime
	rm.RequestTypeId = r.RequestTypeId
	rm.Params = r.Params
	rm.ProcessorId = r.ProcessorId
	rm.RequesterId = r.RequesterId
	rm.WorkflowEnd = r.WorkflowEnd
	rm.BpmnId = r.BpmnId
}
