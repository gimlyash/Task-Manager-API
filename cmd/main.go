package main

import (
	"log"
	"os"

	"TaskManagerApi/internal/handler"
	"github.com/gin-gonic/gin"
	"task-manager-api/internal/repository"
	"task-manager-api/internal/service"
	"task-manager-api/pkg/db"
)

func main() {
	dbConn, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	// Авто миграция схемы
	err = dbConn.AutoMigrate(&repository.TaskRepository{}.Model)
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
