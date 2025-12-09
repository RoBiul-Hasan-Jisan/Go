package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// In-memory storage
var (
	users      = make(map[string]User)
	tasks      = make(map[string][]Task)
	usersMutex sync.RWMutex
	tasksMutex sync.RWMutex
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"due_date"`
	UserID      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required"`
}

func main() {
	// Setup router
	router := gin.Default()
	
	// Allow all CORS for testing
	router.Use(cors.Default())
	
	// Serve frontend files - Use ../frontend since we're running from backend/
	router.Static("/static", "../frontend")
	router.StaticFile("/", "../frontend/index.html")
	router.StaticFile("/login", "../frontend/login.html")
	router.StaticFile("/register", "../frontend/register.html")
	router.StaticFile("/dashboard", "../frontend/dashboard.html")
	
	// API routes
	api := router.Group("/api")
	{
		api.POST("/register", registerHandler)
		api.POST("/login", loginHandler)
		
		// Protected routes
		protected := api.Group("/")
		protected.Use(authMiddleware())
		{
			protected.GET("/tasks", getTasksHandler)
			protected.POST("/tasks", createTaskHandler)
			protected.PUT("/tasks/:id", updateTaskHandler)
			protected.DELETE("/tasks/:id", deleteTaskHandler)
			protected.GET("/user", getUserHandler)
		}
	}
	
	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
			"time":    time.Now().Format(time.RFC3339),
		})
	})
	
	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("🚀 Server starting on http://localhost:%s", port)
	log.Println("📁 Serving frontend from: ../frontend")
	
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}

func registerHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	usersMutex.Lock()
	defer usersMutex.Unlock()
	
	// Check if user exists
	for _, user := range users {
		if user.Email == req.Email || user.Username == req.Username {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}
	}
	
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	
	// Create user
	userID := generateID()
	user := User{
		ID:        userID,
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}
	
	users[userID] = user
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

func loginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	usersMutex.RLock()
	defer usersMutex.RUnlock()
	
	// Find user
	var foundUser *User
	for _, user := range users {
		if user.Email == req.Email {
			foundUser = &user
			break
		}
	}
	
	if foundUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	
	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	
	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": foundUser.ID,
		"email":   foundUser.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":       foundUser.ID,
			"username": foundUser.Username,
			"email":    foundUser.Email,
		},
	})
}

func getTasksHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	
	tasksMutex.RLock()
	userTasks := tasks[userID]
	tasksMutex.RUnlock()
	
	c.JSON(http.StatusOK, userTasks)
}

func createTaskHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	task.ID = generateID()
	task.UserID = userID
	task.CreatedAt = time.Now()
	
	tasksMutex.Lock()
	tasks[userID] = append(tasks[userID], task)
	tasksMutex.Unlock()
	
	c.JSON(http.StatusCreated, task)
}

func updateTaskHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	taskID := c.Param("id")
	
	var updates Task
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	tasksMutex.Lock()
	defer tasksMutex.Unlock()
	
	userTasks, exists := tasks[userID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No tasks found"})
		return
	}
	
	for i, task := range userTasks {
		if task.ID == taskID {
			// Update fields
			if updates.Title != "" {
				userTasks[i].Title = updates.Title
			}
			if updates.Description != "" {
				userTasks[i].Description = updates.Description
			}
			if updates.Status != "" {
				userTasks[i].Status = updates.Status
			}
			if updates.Priority != "" {
				userTasks[i].Priority = updates.Priority
			}
			if !updates.DueDate.IsZero() {
				userTasks[i].DueDate = updates.DueDate
			}
			
			tasks[userID] = userTasks
			c.JSON(http.StatusOK, userTasks[i])
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func deleteTaskHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	taskID := c.Param("id")
	
	tasksMutex.Lock()
	defer tasksMutex.Unlock()
	
	userTasks, exists := tasks[userID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No tasks found"})
		return
	}
	
	for i, task := range userTasks {
		if task.ID == taskID {
			// Remove task
			tasks[userID] = append(userTasks[:i], userTasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func getUserHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	
	usersMutex.RLock()
	user, exists := users[userID]
	usersMutex.RUnlock()
	
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	
	// Hide password
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}
		
		tokenString := authHeader[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})
		
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		
		claims := token.Claims.(jwt.MapClaims)
		c.Set("user_id", claims["user_id"].(string))
		c.Set("user_email", claims["email"].(string))
		
		c.Next()
	}
}

func generateID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(6)
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := range result {
		seed := time.Now().UnixNano() + int64(i)
		result[i] = letters[seed%int64(len(letters))]
	}
	return string(result)
}