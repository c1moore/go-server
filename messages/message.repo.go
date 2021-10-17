package messages

type messageRepo struct {
	messages []Message
}

func newMessageRepo() *messageRepo {
	return &messageRepo{
		messages: make([]Message, 0),
	}
}

func (this *messageRepo) LoadMessages() []Message {
	return this.messages
}

func (this *messageRepo) SaveMessage(m Message) Message {
	this.messages = append(this.messages, m)

	return m
}
