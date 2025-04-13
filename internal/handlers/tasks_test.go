package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"listaPro/internal/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockTaskRepo struct {
	mock.Mock
}

func (m *MockTaskRepo) Create(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepo) GetAllByList(listID uint) ([]models.Task, error) {
	args := m.Called(listID)
	return args.Get(0).([]models.Task), args.Error(1)
}

func TestCreateTaskHandler(t *testing.T) {
	mockRepo := new(MockTaskRepo)
	mockRepo.On("Create", mock.AnythingOfType("*models.Task")).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(
		"POST",
		"/api/lists/1/tasks",
		strings.NewReader(`{"text":"Nova Tarefa"}`),
	)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	CreateTask(mockRepo)(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}
