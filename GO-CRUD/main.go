package main

import (
	"go-basic-crud/handler"
	"go-basic-crud/task"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//define database
	dsn := "host=localhost user=my_gorm_role password=my_gorm_pass dbname=my_gorm_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	//to migrating db schema based on defined entities
	db.AutoMigrate(&task.Task{})

	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()
	api := router.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//task routes
	api.GET("/task", taskHandler.Index)
	api.POST("/task", taskHandler.Store)
	api.GET("/task/:id", taskHandler.Show)
	api.PUT("/task/:id", taskHandler.Update)
	api.DELETE("/task/:id", taskHandler.Destroy)

	router.Run()
}
