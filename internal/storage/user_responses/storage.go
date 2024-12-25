package user_responses

import (
	"Sechenovka/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserResponseStorage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserResponseStorage {
	return &UserResponseStorage{
		db: db,
	}
}

func (s *UserResponseStorage) SaveUserResponse(userId model.UserId, responseId, passNum int) error {
	dal := &UserResponse{
		UserID:     userId.String(),
		ResponseId: responseId,
		PassNum:    passNum,
	}
	if err := s.db.Create(dal).Error; err != nil {
		return errors.Wrap(err, "SaveUserResponse[Storage]")
	}
	return nil
}

//func (s *UserResponseStorage) GetUserTotalScore(userId model.UserID, passNum int) (int, error) {
//	var totalScore int64
//
//	err := s.db.Model(&UserResponse{}).
//		Select("SUM(pass_num)").
//		Where("pass_num = ?", passNum).
//		Where("user_id = ?", userId.String()).
//		Scan(&totalScore).Error
//
//	if err != nil {
//		return 0, err
//	}
//
//	return int(totalScore), nil
//}

func (s *UserResponseStorage) GetUserResponsesByPassNum(userId model.UserId, passNum int) ([]UserResponse, error) {
	userResponses := make([]UserResponse, 0)
	err := s.db.Where("user_id = ?", userId.String()).Where("pass_num = ?", passNum).Order("created_at ASC").Find(&userResponses).Error
	if err != nil {
		return nil, err
	}
	return userResponses, nil
}

func (s *UserResponseStorage) GetLastUserResponse(userId model.UserId) (*UserResponse, error) {
	var userResponse UserResponse
	err := s.db.Where("user_id = ?", userId.String()).Order("created_at DESC").First(&userResponse).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &userResponse, nil
}
