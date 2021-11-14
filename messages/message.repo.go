package messages

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type messageRepo struct {
	collection *mongo.Collection
}

func newMessageRepo(db *mongo.Database) *messageRepo {
	return &messageRepo{
		collection: db.Collection("messages"),
	}
}

func (repo *messageRepo) LoadMessages() (messages []Message, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cursor, err := repo.collection.Find(ctx, bson.D{})
	if err != nil {
		return messages, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &messages)

	if messages == nil {
		messages = make([]Message, 0)
	}

	return messages, err
}

func (repo *messageRepo) SaveMessage(m *Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if m.SentAt.IsZero() {
		m.SentAt = time.Now()
	}

	m.UpdatedAt = m.SentAt

	res, err := repo.collection.InsertOne(ctx, m)
	if err != nil {
		return err
	}

	if id, ok := res.InsertedID.(primitive.ObjectID); ok {
		m.ID = &id
	}

	return err
}
