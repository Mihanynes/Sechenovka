package auth

import (
	"Sechenovka/internal/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type service struct {
	userStorage userStorage
	log         *slog.Logger
}

func New(userStorage userStorage, log *slog.Logger) *service {
	return &service{
		userStorage: userStorage,
		log:         log,
	}
}

func (s *service) Login(snils string, password string) (uuid.UUID, error) {
	userFromDB, err := s.userStorage.GetUser(snils)
	if err != nil {
		return uuid.Nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(password))
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
