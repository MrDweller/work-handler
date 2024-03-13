package workhandler

type WorkerDTO struct {
	WorkerId string `json:"workerId"`
}

func (workerDTO *WorkerDTO) MapToWorker() *Worker {
	return &Worker{
		WorkerId: workerDTO.WorkerId,
	}
}

type Worker struct {
	WorkerId string `json:"workerId"`
}

func (worker *Worker) MapToWorkerDTO() *WorkerDTO {
	return &WorkerDTO{
		WorkerId: worker.WorkerId,
	}
}
