package storage

import (
	"Sechenovka/internal/models"
	"Sechenovka/internal/service/history_saver"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./storage/userAuth.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	//DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	err = db.AutoMigrate(&models.User{}, &history_saver.UserResponse{})
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	log.Println("🚀 Connected Successfully to the Database")
	return db
}
