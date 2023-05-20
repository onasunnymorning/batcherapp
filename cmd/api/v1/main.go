package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	v1 "github.com/onasunnymorning/batcherapp/app/api/v1"
	"github.com/onasunnymorning/batcherapp/batch"
)

type BatchCreateRequest struct {
	Name string `json:"name" binding:"required"` // required field
}

func main() {

	gormDB, err := gorm.Open(postgres.Open("postgres://postgres:mysecretpassword@localhost:5432/batch"))
	if err != nil {
		log.Fatal(err)
	}

	batchRepository := batch.NewPostgreSQLGORMRepository(gormDB)
	batchHandler := v1.NewBatchHandler(batchRepository)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	v1.SetupBatchRoutes(r, batchHandler)

	r.Run(":8080") // listen and serve on port 8080
}
