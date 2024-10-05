package questions

import "Sechenovka/internal/service/history"

type historyStorage interface {
	SaveUserResponse(userResponse *history.UserResponse) error
}
