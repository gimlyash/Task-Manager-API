package service

import (
	"github.com/gimlyash/Task-Manager-API.git/internal/model"
	"github.com/gimlyash/Task-Manager-API.git/internal/repository"
)

type TaskService interface {
	CreateTask(task *model.Task) error
	GetTasks() ([]model.Task, error)
	GetTaskByID(id uint) (*model.Task, error)
	UpdateTask(task *model.Task) error
	DeleteTask(id uint) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task *model.Task) error {
	return s.repo.Create(task)
}

func (s *taskService) GetTasks() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *taskService) GetTaskByID(id uint) (*model.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) UpdateTask(task *model.Task) error {
	return s.repo.Update(task)
}

func (s *taskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
