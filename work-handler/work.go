package workhandler

import "time"

type Work struct {
	WorkId string `json:"workId"`

	ProductId string `json:"productId"`
	EventType string `json:"eventType"`

	Address   string    `json:"address"`
	StartTime time.Time `json:"startTime"`

	IsAssigned       bool   `json:"isAssigned"`
	AssignedWorkerId string `json:"assignedWorkerId"`
}

type CreateWorkDTO struct {
	ProductId string `json:"productId"`
	EventType string `json:"eventType"`
	Address   string `json:"address"`
}

type WorkDTO struct {
	WorkId string `json:"workId"`

	ProductId string `json:"productId"`
	EventType string `json:"eventType"`

	Address   string    `json:"address"`
	StartTime time.Time `json:"startTime"`
}

func (work *Work) MapToWorkDTO() *WorkDTO {
	return &WorkDTO{
		WorkId:    work.WorkId,
		ProductId: work.ProductId,
		EventType: work.EventType,
		Address:   work.Address,
		StartTime: work.StartTime,
	}
}

type AssignWorkerDTO struct {
	WorkerId string `json:"workerId"`
	WorkId   string `json:"workId"`
}
