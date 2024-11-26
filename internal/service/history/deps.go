package history

import history_storage "Sechenovka/internal/storage/user_history"

type historyStorage interface {
	SaveUserResponse(userResponse *history_storage.UserResponse) error
}
