package messages

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Message struct {
	ID *primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	Text string `json:"text" bson:"text"`

	SentAt    time.Time `json:"sentAt" bson:"sent_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}

func Init(engine *gin.Engine, db *mongo.Database) {
	repo := newMessageRepo(db)
	service := newMessageService(repo)
	ctl := newMessageController(service)

	engine.GET("/messages", ctl.GetMessages)

	engine.POST("/messages", ctl.PostMessage)
}
