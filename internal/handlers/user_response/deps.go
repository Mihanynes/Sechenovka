package user_response

import "Sechenovka/internal/model"

type userResponseService interface {
	SaveUserResponse(userResponse *model.UserResponse) (bool, error)
}
