package user_responses

import (
	"Sechenovka/internal/model"
	"github.com/pkg/errors"
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

func (s *storage) SaveUserResponse(userResponse *model.UserResponse) error {
	dal := &UserResponse{
		UserId: userResponse.UserId.String(),
		Response: Response{
			AnswerText: userResponse.Response.AnswerText,
			//AnswerId: userResponse.Response.AnswerId,
			Score: userResponse.Response.Score,
		},
		CorrelationId: userResponse.CorrelationId,
	}
	if err := s.db.Create(dal).Error; err != nil {
		return errors.Wrap(err, "SaveUserResponse[Storage]")
	}
	return nil
}

// GetUserScore Метод для получения суммы score по correlationId
func (s *storage) GetUserTotalScore(userId model.UserId, correlationId string) (int, error) {
	var totalScore int64

	err := s.db.Model(&UserResponse{}).
		Select("SUM(response_score)").
		Where("correlation_id = ?", correlationId).
		Where("user_id = ?", userId.String()).
		Scan(&totalScore).Error

	if err != nil {
		return 0, err
	}

	return int(totalScore), nil
}

func (s *storage) GetUserResponses(userId model.UserId) ([]*model.UserResponse, error) {
	userResponses := make([]UserResponse, 0)
	err := s.db.Where("user_id = ?", userId).Find(&userResponses).Error
	if err != nil {
		return nil, err
	}
	res := make([]*model.UserResponse, len(userResponses))
	for _, userResponse := range userResponses {
		res = append(
			res,
			&model.UserResponse{
				UserId: model.UserIdFromString(userResponse.UserId),
				//QuestionId: userResponse.QuestionId,
				Response: model.Response{
					//AnswerId: userResponse.Response.AnswerId,
					Score: userResponse.Response.Score,
				},
				CorrelationId: "",
			},
		)
	}
	return res, nil
}
