package repository

import (
	"github.com/gimlyash/Task-Manager-API.git/internal/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *model.Task) error
	GetAll() ([]model.Task, error)
	GetByID(id uint) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetAll() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetByID(id uint) (*model.Task, error) {
	var task model.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *taskRepository) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&model.Task{}, id).Error
}
