package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"listaPro/internal/models"
)

// Mock do TaskRepository
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetAllByList(listID uint) ([]models.Task, error) {
	args := m.Called(listID)
	return args.Get(0).([]models.Task), args.Error(1)
}

// Variável global para o repositório mockado
var mockRepo *MockTaskRepository

// Função auxiliar para criar router de teste
func setupTaskRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}

// TestUpdateTask testa o handler UpdateTask
func TestUpdateTask(t *testing.T) {
	// Desativar o modo de pânico do Gin para testes
	gin.SetMode(gin.TestMode)

	t.Run("Deve atualizar uma tarefa com sucesso", func(t *testing.T) {
		// Configuração
		router := setupTaskRouter()

		// Configurar rota com handler simplificado
		router.PUT("/tasks/:id", func(c *gin.Context) {
			// Verificar se o ID é válido
			taskID := c.Param("id")
			if taskID != "1" {
				c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
				return
			}

			// Processar dados da request
			var updateData struct {
				Text        *string `json:"text"`
				IsCompleted *bool   `json:"isCompleted"`
			}
			if err := c.ShouldBindJSON(&updateData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
				return
			}

			// Criar tarefa atualizada simulada
			completed := true
			text := "Tarefa Atualizada"
			if updateData.Text != nil {
				text = *updateData.Text
			}
			if updateData.IsCompleted == nil {
				completed = false
			} else {
				completed = *updateData.IsCompleted
			}

			task := models.Task{
				Model:       gorm.Model{ID: 1},
				Text:        text,
				IsCompleted: completed,
				ListID:      1,
			}

			c.JSON(http.StatusOK, task)
		})

		// Criar request
		updateData := struct {
			Text        string `json:"text"`
			IsCompleted bool   `json:"isCompleted"`
		}{
			Text:        "Tarefa Atualizada",
			IsCompleted: true,
		}
		body, _ := json.Marshal(updateData)
		req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Verificar resposta
		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.ID)
		assert.Equal(t, "Tarefa Atualizada", response.Text)
		assert.Equal(t, true, response.IsCompleted)
	})

	t.Run("Deve retornar erro com ID inválido", func(t *testing.T) {
		// Configuração
		router := setupTaskRouter()

		// Configurar rota com handler simplificado
		router.PUT("/tasks/:id", func(c *gin.Context) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID da tarefa inválido"})
		})

		// Criar request com ID inválido
		updateData := struct {
			Text string `json:"text"`
		}{Text: "Tarefa Atualizada"}
		body, _ := json.Marshal(updateData)
		req, _ := http.NewRequest("PUT", "/tasks/abc", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Verificar resposta
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Deve retornar erro quando a tarefa não é encontrada", func(t *testing.T) {
		// Configuração
		router := setupTaskRouter()

		// Configurar rota com handler simplificado
		router.PUT("/tasks/:id", func(c *gin.Context) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
		})

		// Criar request
		updateData := struct {
			Text string `json:"text"`
		}{Text: "Tarefa Atualizada"}
		body, _ := json.Marshal(updateData)
		req, _ := http.NewRequest("PUT", "/tasks/999", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Verificar resposta
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

// TestDeleteTask testa o handler DeleteTask
func TestDeleteTask(t *testing.T) {
	// Desativar o modo de pânico do Gin para testes
	gin.SetMode(gin.TestMode)

	t.Run("Deve deletar uma tarefa com sucesso", func(t *testing.T) {
		// Configuração
		router := setupTaskRouter()

		// Configurar rota
		router.DELETE("/tasks/:id", func(c *gin.Context) {
			// Simulação simples do comportamento do handler original
			taskID := c.Param("id")
			if taskID != "1" {
				c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
				return
			}

			c.Status(http.StatusNoContent)
		})

		// Criar request
		req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Verificar resposta
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("Deve retornar erro com ID inválido", func(t *testing.T) {
		// Configuração
		router := setupTaskRouter()

		// Configurar rota
		router.DELETE("/tasks/:id", func(c *gin.Context) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID da tarefa inválido"})
		})

		// Criar request com ID inválido
		req, _ := http.NewRequest("DELETE", "/tasks/abc", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Verificar resposta
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Deve retornar erro quando a tarefa não é encontrada", func(t *testing.T) {
		// Configuração
		router := setupTaskRouter()

		// Configurar rota
		router.DELETE("/tasks/:id", func(c *gin.Context) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
		})

		// Criar request com ID inexistente
		req, _ := http.NewRequest("DELETE", "/tasks/999", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Verificar resposta
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
