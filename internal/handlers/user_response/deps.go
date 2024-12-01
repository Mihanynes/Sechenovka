package user_response

import "Sechenovka/internal/model"

type userResponseService interface {
	SaveUserResponse(userResponse *model.UserResponse) (bool, error)
}

type userResponseStorage interface {
	GetUserTotalScore(userId model.UserId, correlationId string) (int, error)
}
