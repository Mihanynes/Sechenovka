package auth

import (
	"Sechenovka/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log/slog"
)

type service struct {
	userStorage userStorage
	log         *slog.Logger
	db          *gorm.DB
}

func New(userStorage userStorage, log *slog.Logger, db *gorm.DB) *service {
	return &service{
		userStorage: userStorage,
		log:         log,
		db:          db,
	}
}

func (s *service) Login(username string, password string) error {
	var userFromDB model.User

	result := s.db.First(&userFromDB, "username = ?", username)

	if result.Error != nil {
		return errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(password))
	if err != nil {
		return errors.New("wrong password")
	}

	return nil
}

func (s *service) Register(user *model.User) error {
	userWithHashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("error while hashing password", err)
		return err
	}
	user.Password = string(userWithHashedPassword)
	generatedUserId := uuid.New()
	return s.userStorage.SaveUser(user, generatedUserId)
}
