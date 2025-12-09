package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getTasks(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var tasks []Task
	if err := db.Where("user_id = ?", userID).Order("created_at DESC").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	
	c.JSON(http.StatusOK, tasks)
}

func getTask(c *gin.Context) {
	userID := c.GetUint("user_id")
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	
	var task Task
	if err := db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		}
		return
	}
	
	c.JSON(http.StatusOK, task)
}

func createTask(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	task.UserID = userID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	
	if err := db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	
	c.JSON(http.StatusCreated, task)
}

func updateTask(c *gin.Context) {
	userID := c.GetUint("user_id")
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	
	var existingTask Task
	if err := db.Where("id = ? AND user_id = ?", taskID, userID).First(&existingTask).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	
	var updates Task
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Update fields
	updates.UpdatedAt = time.Now()
	
	if err := db.Model(&existingTask).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}
	
	c.JSON(http.StatusOK, existingTask)
}

func deleteTask(c *gin.Context) {
	userID := c.GetUint("user_id")
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	
	result := db.Where("id = ? AND user_id = ?", taskID, userID).Delete(&Task{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func getUser(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	
	// Hide password
	user.Password = ""
	c.JSON(http.StatusOK, user)
}