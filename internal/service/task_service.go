package service

import (
	"github.com/gimlyash/Task-Manager-API.git/internal/models"
	"github.com/gimlyash/Task-Manager-API.git/internal/repository"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetTasks() ([]models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id uint) error
}

type taskServiceImpl struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskServiceImpl{repo: repo}
}

func (s *taskServiceImpl) CreateTask(task *models.Task) error {
	return s.repo.Create(task)

}

func (s *taskServiceImpl) GetTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *taskServiceImpl) GetTaskByID(id uint) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskServiceImpl) UpdateTask(task *models.Task) error {
	return s.repo.Update(task)
}

func (s *taskServiceImpl) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
