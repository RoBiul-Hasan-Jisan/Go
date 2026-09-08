package main

import (
	"log"

	"taskmanager/internal/config"
	"taskmanager/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db := config.ConnectDatabase(cfg)

	router := gin.Default()
	routes.Setup(router, db, cfg)

	log.Printf("server starting on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
