package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"listaPro/internal/models"
	"net/http"
	"strconv"
)

func GetAllLists(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var lists []models.TaskList
		if result := db.Preload("Tasks").Find(&lists); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar " +
				"listas"})
			return
		}
		c.JSON(http.StatusOK, lists)
	}
}

func CreateList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newList models.TaskList
		if err := c.ShouldBindJSON(&newList); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao buscar"})
			return
		}

		if result := db.Create(&newList); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar"})
			return
		}

		c.JSON(http.StatusCreated, newList)
	}
}

func UpdateList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao buscar"})
			return
		}

		var updateData struct {
			Name string `json:"name"`
		}
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": "dados inválidos"})
			return
		}

		var list models.TaskList
		result := db.First(&list, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lista não encontrada"})
			return
		}

		list.Name = updateData.Name
		db.Save(&list)

		c.JSON(http.StatusOK, list)
	}
}
//Teste!
func DeleteList(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		result := db.Delete(&models.TaskList{}, id)
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Lista não encontrada"})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
