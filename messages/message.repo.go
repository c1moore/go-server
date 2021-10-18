package messages

import "gorm.io/gorm"

type messageRepo struct {
	db *gorm.DB
}

func newMessageRepo(db *gorm.DB) *messageRepo {
	return &messageRepo{
		db: db,
	}
}

func (repo *messageRepo) LoadMessages() (messages []Message, err error) {
	result := repo.db.Find(&messages)
	if result.Error != nil {
		return messages, result.Error
	}

	return messages, err
}

func (repo *messageRepo) SaveMessage(m *Message) error {
	result := repo.db.Create(m)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
