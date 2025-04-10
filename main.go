package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"listaPro/internal/config"
	"listaPro/internal/handlers"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func main() {

	godotenv.Load()

	db = config.ConnectDB()

	config.Migrate(db)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{

		//listas
		api.GET("/lists", handlers.GetAllLists(db))
		api.POST("/lists", handlers.CreateList(db))
		api.PUT("/lists/:id", handlers.UpdateList(db))
		api.DELETE("/lists/:id", handlers.DeleteList(db))

		//Tasks
		api.GET("/lists/:id/tasks", handlers.GetTasksByList(db))
		api.POST("/lists/:id/tasks", handlers.CreateTask(db))
		api.PUT("/tasks/:id", handlers.UpdateTask(db))
		api.DELETE("/tasks/:id", handlers.DeleteTask(db))
	}

	//Inicia Servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":" + port))
}
