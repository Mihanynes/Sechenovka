package db

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user_history"
	"Sechenovka/internal/storage/users_result"
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
	err = db.AutoMigrate(&model.User{}, &user_history.UserResponse{}, &users_result.UserResult{})
	if err != nil {
		log.Fatal("Migration Failed:\n", err.Error())
		os.Exit(1)
	}

	log.Println("🚀 Connected Successfully to the Database")
	return db
}
