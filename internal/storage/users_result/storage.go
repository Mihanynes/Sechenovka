package users_result

import (
	"Sechenovka/internal/model"
	"context"
	"gorm.io/gorm"
)

type storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *storage {
	return &storage{
		db: db,
	}
}

func (s *storage) Save(ctx context.Context, userResult *model.UserResult) error {

}

func (s *storage) GetByUserIds(ctx context.Context, userIds []model.UserId) ([]*model.UserResult, error) {

}
