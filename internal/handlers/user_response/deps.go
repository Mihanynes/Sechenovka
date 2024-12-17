package user_response

import "Sechenovka/internal/model"

type userResponseService interface {
	SaveUserResponse(userId model.UserId, responseId, passNum int) (bool, error)
}
