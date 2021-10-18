package messages

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`

	Text string `json:"text"`

	SentAt    time.Time `json:"sentAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (message *Message) BeforeCreate(db *gorm.DB) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	message.ID = uuid

	return nil
}

func Init(engine *gin.Engine, db *gorm.DB) {
	repo := newMessageRepo(db)
	service := newMessageService(repo)
	ctl := newMessageController(service)

	engine.GET("/messages", ctl.GetMessages)

	engine.POST("/messages", ctl.PostMessage)
}
