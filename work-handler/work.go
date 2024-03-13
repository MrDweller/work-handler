package workhandler

type Work struct {
	WorkId string `json:"workId"`

	ProductId string `json:"productId"`
	EventType string `json:"eventType"`

	IsAssigned       bool   `json:"isAssigned"`
	AssignedWorkerId string `json:"assignedWorkerId"`
}

type CreateWorkDTO struct {
	ProductId string `json:"productId"`
	EventType string `json:"eventType"`
}

type WorkDTO struct {
	WorkId string `json:"workId"`

	ProductId string `json:"productId"`
	EventType string `json:"eventType"`
}

func (work *Work) MapToWorkDTO() *WorkDTO {
	return &WorkDTO{
		WorkId:    work.WorkId,
		ProductId: work.ProductId,
		EventType: work.EventType,
	}
}

type AssignWorkerDTO struct {
	WorkerId string `json:"workerId"`
	WorkId   string `json:"workId"`
}
