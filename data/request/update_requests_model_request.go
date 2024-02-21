package request

import (
	"golang-crud-gin/model"
	"time"
)

type UpdateRequestsModelRequest struct {
	Id            int64     `json:"id"`
	CurrentNode   string    `json:"current_node"`
	Status        string    `json:"status"`
	CreateTime    time.Time `json:"create_time"`
	FinishTime    time.Time `json:"finish_time"`
	RequestTypeId int       `json:"request_type_id omitempty"`
	Params        string    `json:"params omitempty"`
	ProcessorId   *int      `json:"processor_id omitempty"`
	RequesterId   *int      `json:"requester_id omitempty"`
	WorkflowEnd   byte      `json:"workflow_end omitempty"`
	BpmnId        int       `json:"bpmn_id omitempty"`
}

func UpdateRequestConvertToModel(r *UpdateRequestsModelRequest, rm *model.RequestsModel) {
	rm.Id = r.Id
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
