package main

import (
	"github.com/gimlyash/Task-Manager-API.git/internal/database"
	"github.com/gimlyash/Task-Manager-API.git/internal/handler"
	"github.com/gimlyash/Task-Manager-API.git/internal/model"
	"github.com/gimlyash/Task-Manager-API.git/internal/repository"
	"github.com/gimlyash/Task-Manager-API.git/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	dbConn, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	err = dbConn.AutoMigrate(&model.Task{})
	if err != nil {
		log.Fatalf("failed migration: %v", err)
	}

	taskRepo := repository.NewTaskRepository(dbConn)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	r := gin.Default()

	tasks := r.Group("/tasks")
	{
		tasks.POST("", taskHandler.CreateTask)
		tasks.GET("", taskHandler.GetTasks)
		tasks.GET("/:id", taskHandler.GetTaskByID)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.DeleteTask)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
