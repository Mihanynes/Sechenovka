package user_response

import "Sechenovka/internal/model"

type userResponseService interface {
	SaveUserResponses(userId model.UserId, responseIds []int, passNum int) (bool, error)
}
