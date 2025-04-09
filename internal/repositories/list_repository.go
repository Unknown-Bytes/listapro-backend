package repositories

import (
	"gorm.io/gorm"
	"listaPro/internal/models"
)

type ListRepository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) *ListRepository {
	return &ListRepository{db: db}
}

// Create - Cria uma nova lista
func (r *ListRepository) Create(list *models.TaskList) error {
	return r.db.Create(list).Error
}

// GetAll - Retorna todas as listas com suas tarefas
func (r *ListRepository) GetAll() ([]models.TaskList, error) {
	var lists []models.TaskList
	err := r.db.Preload("Tasks").Find(&lists).Error
	return lists, err
}

// GetByID - Busca uma lista por ID com suas tarefas
func (r *ListRepository) GetByID(id uint) (*models.TaskList, error) {
	var list models.TaskList
	err := r.db.Preload("Tasks").First(&list, id).Error
	return &list, err
}

// Update - Atualiza o nome de uma lista
func (r *ListRepository) Update(list *models.TaskList) error {
	return r.db.Model(list).Update("name", list.Name).Error
}

// Delete - Exclui uma lista por ID (CASCADE via GORM)
func (r *ListRepository) Delete(id uint) error {
	return r.db.Delete(&models.TaskList{}, id).Error
}

// Exists - Verifica se uma lista existe
func (r *ListRepository) Exists(id uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.TaskList{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}
