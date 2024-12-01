package patient_info

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user"
)

type userStorage interface {
	GetUserByUserId(userId model.UserId) (*user.User, error)
}
