package config

import (
	"fmt"
	"listaPro/internal/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	//Formatar DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	// Conexão com banco
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados!")
	}
	return db
}

func Migrate(db *gorm.DB) {
	// Executar migrações
	err := db.AutoMigrate(&models.TaskList{}, &models.Task{})
	if err != nil {
		panic("Falha ao migrar tabelas: " + err.Error())
	}
}
