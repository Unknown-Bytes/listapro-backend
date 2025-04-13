package repositories

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"listaPro/internal/models"
	"testing"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{}) // Configuração correta
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.TaskList{}, &models.Task{})
	assert.NoError(t, err)
	return db
}

func TestCreateList(t *testing.T) {
	db := setupTestDB(t)
	repo := NewListRepository(db)

	list := &models.TaskList{Name: "Compras"}
	err := repo.Create(list)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), list.ID) // Verifica se o ID foi gerado
}
