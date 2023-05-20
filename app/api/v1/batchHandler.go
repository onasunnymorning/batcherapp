package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/onasunnymorning/batcherapp/batch"
)

type BatchCreateRequest struct {
	Name string `json:"name" binding:"required"` // required field
}

type BatchNotesRequest struct {
	Notes string `json:"notes" binding:"required"` // required field
}

type batchHandler struct {
	batchRepository batch.Repository
}

func NewBatchHandler(batchRepository batch.Repository) *batchHandler {
	return &batchHandler{
		batchRepository: batchRepository,
	}
}

func (h *batchHandler) CreateBatch(c *gin.Context) {
	var newBatchRequest BatchCreateRequest
	if err := c.ShouldBindJSON(&newBatchRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newBatch := batch.NewBatch(newBatchRequest.Name)
	createdBatch, err := h.batchRepository.Create(c, newBatch)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, createdBatch)
}

func (h *batchHandler) GetBatchByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	batch, err := h.batchRepository.GetByID(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, batch)
}

func (h *batchHandler) GetBatchByName(c *gin.Context) {
	name := c.Param("name")
	batch, err := h.batchRepository.GetByName(c, name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, batch)
}

func (h *batchHandler) DeleteBatchByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = h.batchRepository.DeleteByID(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "batch deleted"})
}

func (h *batchHandler) AddIngredientToBatch(c *gin.Context) {
	var ingredient batch.BatchIngredient
	if err := c.ShouldBindJSON(&ingredient); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	batch, err := h.batchRepository.GetByID(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	batch.AddIngredient(ingredient)
	updatedBatch, err := h.batchRepository.Update(c, *batch)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedBatch)
}

func (h *batchHandler) UpdateBatchNotes(c *gin.Context) {
	var notes BatchNotesRequest
	if err := c.ShouldBindJSON(&notes); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	batch, err := h.batchRepository.GetByID(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	batch.Notes = notes.Notes
	updatedBatch, err := h.batchRepository.Update(c, *batch)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedBatch)
}

func (h *batchHandler) StopBatch(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	batch, err := h.batchRepository.GetByID(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	batch.Stop()
	updatedBatch, err := h.batchRepository.Update(c, *batch)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedBatch)
}

func (h *batchHandler) MarkAptForProduction(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	batch, err := h.batchRepository.GetByID(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	batch.MarkAptForProduction()
	updatedBatch, err := h.batchRepository.Update(c, *batch)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedBatch)
}

func (h *batchHandler) RemoveAptForProduction(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	batch, err := h.batchRepository.GetByID(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	batch.RemoveAptForProduction()
	updatedBatch, err := h.batchRepository.Update(c, *batch)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedBatch)
}
