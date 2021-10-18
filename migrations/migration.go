package migrations

import (
	"github.com/c1moore/go-server/messages"
	"gorm.io/gorm"
)

func RunDBMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&messages.Message{})
}
