package apiv1

import (
	"github.com/gin-gonic/gin"
)

func SetupBatchRoutes(r *gin.Engine, h *batchHandler) {
	h.batchRepository.Migrate(nil)
	r.POST("/batch", h.CreateBatch)
	r.GET("/batch/:id", h.GetBatchByID)
	r.DELETE("/batch/:id", h.DeleteBatchByID)
	r.GET("/batch/name/:name", h.GetBatchByName)

	r.POST("/batch/:id/ingredient", h.AddIngredientToBatch)
	r.POST("/batch/:id/notes", h.UpdateBatchNotes)
	r.POST("/batch/:id/stop", h.StopBatch)
	r.POST("/batch/:id/production", h.MarkAptForProduction)
	r.DELETE("/batch/:id/production", h.RemoveAptForProduction)
}
