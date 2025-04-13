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

type MockListRepository struct {
	mock.Mock
}

func (m *MockListRepository) Create(list *models.TaskList) error {
	args := m.Called(list)
	return args.Error(0)
}

func TestCreateListHandler(t *testing.T) {
	mockRepo := new(MockListRepository)
	mockRepo.On("Create", mock.AnythingOfType("*models.TaskList")).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/api/lists", strings.NewReader(`{"name":"Test List"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	CreateList(mockRepo)(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}
