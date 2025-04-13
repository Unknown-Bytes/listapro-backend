package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"listaPro/internal/handlers"
	"listaPro/internal/models"
	"listaPro/internal/repositories"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// setupTestRouter configura um router Gin para testes
func setupTestRouter() *gin.Engine {
	// Banco de dados em memória
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrações
	db.AutoMigrate(&models.TaskList{}, &models.Task{})

	// Repositórios
	listRepo := repositories.NewListRepository(db)

	// Router Gin
	r := gin.Default()

	// Rotas (simulando seu main.go)
	r.GET("/api/lists", handlers.GetAllLists(listRepo))
	r.POST("/api/lists", handlers.CreateList(listRepo))

	return r
}

// TestGetLists verifica o endpoint GET /api/lists
func TestGetLists(t *testing.T) {
	// Configuração
	router := setupTestRouter()

	// Cria dados de teste diretamente no banco (opcional)
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.Create(&models.TaskList{Name: "Lista de Teste"})

	// Requisição
	req, _ := http.NewRequest("GET", "/api/lists", nil)
	w := httptest.NewRecorder()

	// Execução
	router.ServeHTTP(w, req)

	// Verificações
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Lista de Teste")
}

// TestCreateList verifica o endpoint POST /api/lists
func TestCreateList(t *testing.T) {
	router := setupTestRouter()

	// JSON de exemplo
	jsonBody := `{"name": "Nova Lista"}`

	// Requisição
	req, _ := http.NewRequest("POST", "/api/lists", strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Execução
	router.ServeHTTP(w, req)

	// Verificações
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Nova Lista")
}
