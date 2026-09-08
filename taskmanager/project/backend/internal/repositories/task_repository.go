package repositories

import (
	"taskmanager/internal/models"

	"gorm.io/gorm"
)

// TaskRepository handles all database access for tasks.
type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.DB.Create(task).Error
}

// FindAllByUser returns every task belonging to a given user, newest first.
func (r *TaskRepository) FindAllByUser(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// FindByIDAndUser fetches a task, scoped to its owner, so users can never
// read or modify another user's tasks.
func (r *TaskRepository) FindByIDAndUser(id, userID uint) (*models.Task, error) {
	var task models.Task
	if err := r.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Update(task *models.Task) error {
	return r.DB.Save(task).Error
}

func (r *TaskRepository) Delete(task *models.Task) error {
	return r.DB.Delete(task).Error
}
