package repositories

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"listaPro/internal/models"
	"testing"
)

func setupTaskTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&models.TaskList{}, &models.Task{})
	return db
}

func TestCreateTask(t *testing.T) {
	db := setupTaskTestDB(t)
	repo := NewTaskRepository(db)

	// Cria uma lista pai
	list := &models.TaskList{Name: "Lista de Teste"}
	db.Create(list)

	task := &models.Task{
		Text:   "Tarefa de Teste",
		ListID: list.ID,
	}
	err := repo.Create(task)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), task.ID)
}

func TestGetTasksByList(t *testing.T) {
	db := setupTaskTestDB(t)
	repo := NewTaskRepository(db)

	// Setup: Cria lista e tarefas
	list := &models.TaskList{Name: "Lista com Tarefas"}
	db.Create(list)
	db.Create(&models.Task{Text: "Tarefa 1", ListID: list.ID})
	db.Create(&models.Task{Text: "Tarefa 2", ListID: list.ID})

	tasks, err := repo.GetAllByList(list.ID)

	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
}
