package db

import (
	"Sechenovka/internal/storage/doctor_patient"
	"Sechenovka/internal/storage/user"
	"Sechenovka/internal/storage/user_responses"
	"Sechenovka/internal/storage/user_result"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const path = "/app/data/master.db"

func ConnectDB() *gorm.DB {
	dbPath := path
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	//DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.Logger = logger.Default.LogMode(logger.Error)

	log.Println("Running Migrations")
	err = db.AutoMigrate(&user.User{}, &user_responses.UserResponse{}, &user_result.UserResult{}, &doctor_patient.DoctorPatient{})
	if err != nil {
		log.Fatal("Migration Failed:\n", err.Error())
		os.Exit(1)
	}

	log.Println("🚀 Connected Successfully to the Database")
	return db
}
