package messages

type messageService struct {
	repo *messageRepo
}

func newMessageService(r *messageRepo) *messageService {
	return &messageService{
		repo: r,
	}
}

func (this *messageService) LoadMessages() []Message {
	return this.repo.LoadMessages()
}

func (this *messageService) AddMessage(m Message) Message {
	return this.repo.SaveMessage(m)
}
