package workhandler

import (
	"context"

	"github.com/MrDweller/work-handler/database"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type ServiceImplementation struct {
}

func (s *ServiceImplementation) CreateWork(createWorkDTO *CreateWorkDTO) (*WorkDTO, error) {
	work := &Work{
		WorkId: uuid.New().String(),

		ProductId: createWorkDTO.ProductId,
		EventType: createWorkDTO.EventType,

		IsAssigned:       false,
		AssignedWorkerId: "",
	}

	_, err := database.Work.InsertOne(context.Background(), work)
	if err != nil {
		return nil, err
	}
	return work.MapToWorkDTO(), nil
}

func (s *ServiceImplementation) HasWorker(workId string) (bool, error) {
	filter := bson.M{
		"workid": workId,
	}
	var work Work
	err := database.Work.FindOne(context.Background(), filter).Decode(&work)
	if err != nil {
		return true, err
	}

	return work.IsAssigned, nil
}

func (s *ServiceImplementation) IsWorker(workId string, workerId string) (bool, error) {
	filter := bson.M{
		"workid": workId,
	}
	var work Work
	err := database.Work.FindOne(context.Background(), filter).Decode(&work)
	if err != nil {
		return false, err
	}

	if work.AssignedWorkerId != workerId {
		return false, nil
	}
	return true, nil
}

func (s *ServiceImplementation) AssignWorker(assignWorkerDTO *AssignWorkerDTO) (*WorkDTO, error) {
	filter := bson.M{
		"workid": assignWorkerDTO.WorkId,
	}
	update := bson.M{
		"$set": bson.M{
			"isassigned":       true,
			"assignedworkerid": assignWorkerDTO.WorkerId,
		},
	}

	_, err := database.Work.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	filter = bson.M{
		"workid": assignWorkerDTO.WorkId,
	}
	var work Work
	err = database.Work.FindOne(context.Background(), filter).Decode(&work)
	if err != nil {
		return nil, err
	}

	return work.MapToWorkDTO(), nil
}
