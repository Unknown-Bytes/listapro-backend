package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"listaPro/internal/models"
	"listaPro/internal/repositories"
	"net/http"
	"strconv"
)

// GetTasksByList - Obter todas as tarefas de uma lista específica
func GetTasksByList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		listID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		repo := repositories.NewTaskRepository(db)

		tasks, err := repo.GetAllByList(uint(listID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tarefas"})
			return
		}

		c.JSON(http.StatusOK, tasks)
	}
}

// CreateTask (POST /api/lists/:id/tasks)
func CreateTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		listID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": "ID da lista inválido"})
			return
		}

		var taskData struct {
			Text string `json:"text"`
		}
		if err := c.ShouldBindJSON(&taskData); err != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": "Dados inválidos"})
			return
		}

		task := models.Task{
			Text:   taskData.Text,
			ListID: uint(listID),
		}

		if result := db.Create(&task); result.Error != nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{"error": "Erro ao crear task"})
			return
		}

		c.JSON(http.StatusCreated, task)
	}
}

// UpdateTask (PUT /api/tasks/:id)
func UpdateTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID da tarefa inválido"})
			return
		}

		var updateData struct {
			Text        *string `json:"text"`
			IsCompleted *bool   `json:"isCompleted"`
		}
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}

		var task models.Task
		if result := db.First(&task, taskID); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
			return
		}

		// atualiza campos fornecidos
		if updateData.Text != nil {
			task.Text = *updateData.Text
		}
		if updateData.IsCompleted != nil {
			task.IsCompleted = *updateData.IsCompleted
		}

		db.Save(&task)

		c.JSON(http.StatusOK, task)
	}
}

// DeleteTask (DELETE /api/tasks/:id)
func DeleteTask(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID da tarefa inválido"})
			return
		}

		result := db.Delete(&models.Task{}, taskID)
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
