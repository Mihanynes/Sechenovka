package auth

import (
	"Sechenovka/internal/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log/slog"
)

type service struct {
	log *slog.Logger
	db  *gorm.DB
}

func New(log *slog.Logger, db *gorm.DB) *service {
	return &service{
		log: log,
		db:  db,
	}
}

func (s *service) Login(username string, password string) error {
	var userFromDB models.User

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

func (s *service) Register(user *models.User) error {
	userWithHashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("error while hashing password", err)
		return err
	}
	user.Password = string(userWithHashedPassword)

	result := s.db.Create(&user)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		s.log.Warn("user already exists")
		return errors.New("user already exists")
	}
	if result.Error != nil {
		s.log.Error("saving user error", result.Error)
		return result.Error
	}
	return nil
}
