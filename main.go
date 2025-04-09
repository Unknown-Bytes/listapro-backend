package main

import (
    "gorm.io/gorm"
    "listaPro/internal/config"
    "listaPro/internal/handlers"
    "log"
    "os"

  //"github.com/gin-gonic/gin"
  //"gorm.io.io/gorm"
)

var db *gorm.DB

func main() {

    db = config.ConnectDB()

    config.Migrate(db)

    router := gin.Default()

    api := router.Group("/api"){

       //listas
       api.GET("/lists", handlers.GetAllLists(db))
       api.POST("/lists", handlers.CreateList(db))
       api.PUT("/lists/:id", handlers.UpdateList(db))
       api.DELETE("/lists/:id", handlers.DeleteList(db))

       //Taskss
       api.GET("/lists/:id/tasks", handlers.GetTaksByList(db))
       api.POST("/lists/:id/tasks", handlers.CreateTask(db))
       api.PUT("/lists/:id", handlers.UpdateTask(db))
       api.DELETE("/lists/:id", handlers.DeleteTask(db))
  }

  //Inicia Servidor
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }
  log.Fatal(router.Run(":" + port))
}
