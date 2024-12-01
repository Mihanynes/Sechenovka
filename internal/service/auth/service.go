package auth

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user_info"
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

func (s *service) Login(snils string, password string) (uuid.UUID, error) {
	var userFromDB user_info.User

	result := s.db.First(&userFromDB, "snils = ?", snils)

	if result.Error != nil {
		return uuid.Nil, errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(password))
	if err != nil {
		return uuid.Nil, errors.New("wrong password")
	}

	return userFromDB.UserId, nil
}

func (s *service) Register(user *model.User) error {
	userWithHashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("error while hashing password", err)
		return err
	}
	user.Password = string(userWithHashedPassword)
	generatedUserId := uuid.New()
	if err := s.userStorage.SaveUser(user, generatedUserId); err != nil {
		return err
	}
	return nil
}
