package services

import (
	"errors"

	"taskmanager/internal/models"
	"taskmanager/internal/repositories"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskService struct {
	taskRepo *repositories.TaskRepository
}

func NewTaskService(taskRepo *repositories.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) CreateTask(userID uint, title, description string) (*models.Task, error) {
	task := &models.Task{
		UserID:      userID,
		Title:       title,
		Description: description,
		Completed:   false,
	}
	if err := s.taskRepo.Create(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) GetTasks(userID uint) ([]models.Task, error) {
	return s.taskRepo.FindAllByUser(userID)
}

func (s *TaskService) GetTask(id, userID uint) (*models.Task, error) {
	task, err := s.taskRepo.FindByIDAndUser(id, userID)
	if err != nil {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

// UpdateTask applies partial updates to a task. Pointer fields let the
// handler distinguish between "not provided" and "set to zero value".
func (s *TaskService) UpdateTask(id, userID uint, title, description *string, completed *bool) (*models.Task, error) {
	task, err := s.taskRepo.FindByIDAndUser(id, userID)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	if title != nil {
		task.Title = *title
	}
	if description != nil {
		task.Description = *description
	}
	if completed != nil {
		task.Completed = *completed
	}

	if err := s.taskRepo.Update(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) DeleteTask(id, userID uint) error {
	task, err := s.taskRepo.FindByIDAndUser(id, userID)
	if err != nil {
		return ErrTaskNotFound
	}
	return s.taskRepo.Delete(task)
}
