package db

import (
	"log"

	"todo-server/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() {

	const dsn = "host=localhost port=5432 user=postgres password=password dbname=todo_app sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Could not connect to db")
	}

	log.Println("Database connection is successful")

	db.AutoMigrate(&model.Todo{})
	DB = db

}
