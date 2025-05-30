package main

import (
	"Task-Manager-API/pkg/db"
	"github.com/gimlyash/Task-Manager-API.git/internal/models"
	"github.com/gimlyash/Task-Manager-API.git/internal/repository"
	"github.com/gimlyash/Task-Manager-API.git/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	// Авто миграция схемы
	err = dbConn.AutoMigrate(&models.Task{})
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
