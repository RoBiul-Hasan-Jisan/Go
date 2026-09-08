package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"taskmanager/internal/middleware"
	"taskmanager/internal/services"
	"taskmanager/internal/utils"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

type createTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type updateTaskRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

// getTaskID parses and validates the :id URL param.
func getTaskID(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// CreateTask handles POST /api/tasks
func (h *TaskHandler) CreateTask(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req createTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	task, err := h.taskService.CreateTask(userID, req.Title, req.Description)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to create task")
		return
	}

	utils.Success(c, http.StatusCreated, "Task created successfully", task)
}

// GetTasks handles GET /api/tasks
func (h *TaskHandler) GetTasks(c *gin.Context) {
	userID := middleware.GetUserID(c)

	tasks, err := h.taskService.GetTasks(userID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to fetch tasks")
		return
	}

	utils.Success(c, http.StatusOK, "Tasks retrieved successfully", tasks)
}

// GetTask handles GET /api/tasks/:id
func (h *TaskHandler) GetTask(c *gin.Context) {
	userID := middleware.GetUserID(c)

	id, err := getTaskID(c)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, err := h.taskService.GetTask(id, userID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Task not found")
		return
	}

	utils.Success(c, http.StatusOK, "Task retrieved successfully", task)
}

// UpdateTask handles PUT /api/tasks/:id
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	userID := middleware.GetUserID(c)

	id, err := getTaskID(c)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	var req updateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	task, err := h.taskService.UpdateTask(id, userID, req.Title, req.Description, req.Completed)
	if err != nil {
		if errors.Is(err, services.ErrTaskNotFound) {
			utils.Error(c, http.StatusNotFound, "Task not found")
			return
		}
		utils.Error(c, http.StatusInternalServerError, "Failed to update task")
		return
	}

	utils.Success(c, http.StatusOK, "Task updated successfully", task)
}

// DeleteTask handles DELETE /api/tasks/:id
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	userID := middleware.GetUserID(c)

	id, err := getTaskID(c)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid task ID")
		return
	}

	if err := h.taskService.DeleteTask(id, userID); err != nil {
		if errors.Is(err, services.ErrTaskNotFound) {
			utils.Error(c, http.StatusNotFound, "Task not found")
			return
		}
		utils.Error(c, http.StatusInternalServerError, "Failed to delete task")
		return
	}

	utils.Success(c, http.StatusOK, "Task deleted successfully", nil)
}
