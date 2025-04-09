package config

import (
	"fmt"
	"listaPro/internal/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados: " + err.Error())
	}

	return db
}

func Migrate(db *gorm.DB) {
	// Executar migrações
	fmt.Println("Migrating database ...")
	err := db.AutoMigrate(&models.TaskList{}, &models.Task{})
	if err != nil {
		panic("Falha ao migrar tabelas: " + err.Error())
	}
}
