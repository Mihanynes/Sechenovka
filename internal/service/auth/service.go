package auth

import (
	"Sechenovka/internal/handlers/auth"
	"Sechenovka/internal/model"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log/slog"
)

var ErrUserAlreadyExists = errors.New("User already exists")

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

func (s *service) Login(username string, password string) (*auth.LoginOut, error) {
	userFromDB, err := s.userStorage.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(password))
	if err != nil {
		return nil, errors.New("wrong password")
	}

	return &auth.LoginOut{
		UserId:  userFromDB.UserID,
		IsAdmin: userFromDB.IsAdmin,
	}, nil
}

func (s *service) Register(user *model.User) (*model.UserId, error) {
	userWithHashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("error while hashing password %v", err.Error())
		return nil, err
	}
	user.Password = string(userWithHashedPassword)
	generatedUserId := model.UserId(uuid.New())

	err = s.userStorage.SaveUser(user, generatedUserId)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, ErrUserAlreadyExists
	}
	return &generatedUserId, nil
}
