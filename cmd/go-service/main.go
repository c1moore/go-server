package main

import (
	"net/http"

	"github.com/c1moore/go-server/messages"
	"github.com/c1moore/go-server/migrations"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	envVars := InitEnvVars()

	db, err := gorm.Open(postgres.Open(envVars.Database.ConnectionUri), &gorm.Config{})
	if err != nil {
		panic("Could not connect to DB: " + err.Error())
	}

	if err := migrations.RunDBMigrations(db); err != nil {
		panic("Could not run DB migrations: " + err.Error())
	}

	engine := gin.Default()

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"healthy": true,
		})
	})

	messages.Init(engine, db)

	engine.Run()
}
