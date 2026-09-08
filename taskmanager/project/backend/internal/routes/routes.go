package routes

import (
	"taskmanager/internal/config"
	"taskmanager/internal/handlers"
	"taskmanager/internal/middleware"
	"taskmanager/internal/repositories"
	"taskmanager/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Setup wires up repositories, services, handlers, and registers all routes.
func Setup(router *gin.Engine, db *gorm.DB, cfg *config.Config) {
	router.Use(middleware.CORS(cfg.FrontendURL))

	// Repositories
	userRepo := repositories.NewUserRepository(db)
	taskRepo := repositories.NewTaskRepository(db)

	// Services
	authService := services.NewAuthService(userRepo, cfg)
	taskService := services.NewTaskService(taskRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)

	api := router.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)

		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		tasks := api.Group("/tasks")
		tasks.Use(middleware.AuthRequired(cfg.JWTSecret))
		{
			tasks.GET("", taskHandler.GetTasks)
			tasks.POST("", taskHandler.CreateTask)
			tasks.GET("/:id", taskHandler.GetTask)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}
	}
}
