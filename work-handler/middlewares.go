package workhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CheckWorkerDTO struct {
	WorkId   string `json:"workId"`
	WorkerId string `json:"workerId"`
}

func (controller *Controller) CheckCurrentWorker(c *gin.Context) {
	var checkWorkerDTO CheckWorkerDTO
	if err := c.ShouldBindBodyWith(&checkWorkerDTO, binding.JSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	hasWorker, err := controller.service.HasWorker(checkWorkerDTO.WorkId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	isWorker, err := controller.service.IsWorker(checkWorkerDTO.WorkId, checkWorkerDTO.WorkerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	if hasWorker && !isWorker {
		c.JSON(http.StatusBadRequest, gin.H{"error": "that work task is occupied by another worker"})
		c.Abort()
		return
	}

}
