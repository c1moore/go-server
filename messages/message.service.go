package messages

type messageService struct {
	repo *messageRepo
}

func newMessageService(r *messageRepo) *messageService {
	return &messageService{
		repo: r,
	}
}

func (srv *messageService) LoadMessages() ([]Message, error) {
	return srv.repo.LoadMessages()
}

func (srv *messageService) AddMessage(m *Message) error {
	return srv.repo.SaveMessage(m)
}
