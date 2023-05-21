package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	v1 "github.com/onasunnymorning/batcherapp/app/api/v1"
	"github.com/onasunnymorning/batcherapp/batch"
)

type BatchCreateRequest struct {
	Name string `json:"name" binding:"required"` // required field
}

func main() {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// gormDB, err := gorm.Open(postgres.Open("postgres://postgres:mysecretpassword@batcherapp-db-1:5432/batch"))
	gormDB, err := gorm.Open(postgres.Open("postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName))
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
