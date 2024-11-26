package history

type service struct {
	storage historyStorage
}

func NewSaver(storage historyStorage) *service {
	return &service{
		storage: storage,
	}
}
