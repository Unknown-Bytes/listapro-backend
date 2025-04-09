package repositories

import (
	"gorm.io/gorm"
	"listaPro/internal/models"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// Create cria uma nova tarefa
func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

// GetByID busca uma tarefa pelo ID
func (r *TaskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

// GetAllByList busca todas as tarefas de uma lista
func (r *TaskRepository) GetAllByList(listID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("list_id = ?", listID).Find(&tasks).Error
	return tasks, err
}

// Update atualiza uma tarefa
func (r *TaskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

// Delete remove uma tarefa
func (r *TaskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}

// MarkAsCompleted marca uma tarefa como concluída
func (r *TaskRepository) MarkAsCompleted(id uint) error {
	return r.db.Model(&models.Task{}).
		Where("id = ?", id).
		Update("is_completed", true).Error
}
