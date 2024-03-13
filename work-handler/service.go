package workhandler

type Service interface {
	CreateWork(createWorkDTO *CreateWorkDTO) (*WorkDTO, error)
	HasWorker(workId string) (bool, error)
	IsWorker(workId string, workerId string) (bool, error)
	AssignWorker(assignWorkerDTO *AssignWorkerDTO) (*WorkDTO, error)
}
