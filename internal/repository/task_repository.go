package repository

import (
	"github.com/gimlyash/Task-Manager-API.git/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *models.Task) error
	GetAll() ([]models.Task, error)
	GetByID(id uint) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
}

type taskRepositoryImpl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepositoryImpl{db: db}
}

func (r *taskRepositoryImpl) Create(task *models.Task) error {
	return r.db.Create(&task).Error
}

func (r *taskRepositoryImpl) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepositoryImpl) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, "id = ?", id).Error
	return &task, err
}

func (r *taskRepositoryImpl) Update(task *models.Task) error {
	return r.db.Save(&task).Error
}

func (r *taskRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, "id = ?", id).Error
}
