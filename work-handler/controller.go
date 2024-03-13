package workhandler

import (
	"net/http"

	_ "github.com/MrDweller/work-handler/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Controller struct {
	service Service
}

// Create a new work task
// @Summary      Create a new work task
// @Description  Create a new work task
// @Tags         Product
// @Produce      json
// @Param        CreateWorkDTO  body       CreateWorkDTO  true  "CreateWorkDTO JSON"
// @Success      200 {object} WorkDTO
// @Router       /work [post]
func (controller *Controller) CreateWork(c *gin.Context) {
	var createWorkDTO CreateWorkDTO
	if err := c.BindJSON(&createWorkDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	workDTO, err := controller.service.CreateWork(&createWorkDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workDTO)
}

// Assign a worker a specified work task.
// @Summary      Assign a worker a specified work task.
// @Description  Assign a worker a specified work task.
// @Tags         Worker
// @Produce      json
// @Param        AssignWorkerDTO  body       AssignWorkerDTO  true  "AssignWorkerDTO JSON"
// @Success      200 {object} WorkDTO
// @Router       /assign-worker [post]
func (controller *Controller) AssignWorker(c *gin.Context) {
	var assigneWorkerDTO AssignWorkerDTO
	if err := c.ShouldBindBodyWith(&assigneWorkerDTO, binding.JSON); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	workerDTO, err := controller.service.AssignWorker(&assigneWorkerDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workerDTO)

}
