package history

import (
	"Sechenovka/internal/queue"
	"fmt"
)

type saver struct {
	storage *historyStorage
}

func NewSaver(storage *historyStorage) *saver {
	return &saver{
		storage: storage,
	}
}

func (s *saver) Process(message queue.Message) error {
	userResponse, ok := message.(*UserResponse)
	if !ok {
		return fmt.Errorf("unexpected message type: %T", message)
	}
	err := s.storage.SaveUserResponse(userResponse)
	if err != nil {
		return err
	}
	return nil
}
