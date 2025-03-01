package auth

import (
	"Sechenovka/internal/model"
	"fmt"
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

func (s *service) Login(username string, password string) (string, error) {
	userFromDB, err := s.userStorage.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	return userFromDB.UserID, nil
}

func (s *service) Register(user *model.User) (*model.UserId, error) {
	userWithHashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("error while hashing password %v", err.Error())
		return nil, err
	}
	user.Password = string(userWithHashedPassword)
	generatedUserId := model.UserId(uuid.New())

	if err := s.userStorage.SaveUser(user, generatedUserId); err != nil {
		err = fmt.Errorf("error while saving user %v", err.Error())
		return nil, err
	}
	return &generatedUserId, nil
}
