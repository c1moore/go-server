package main

import (
	"context"
	"net/http"
	"time"

	"github.com/c1moore/go-server/messages"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	envVars := InitEnvVars()

	client, err := connectToDatabase(envVars)
	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	defer func() {
		ctx := context.Background()

		client.Disconnect(ctx)
	}()

	engine := gin.Default()

	engine.GET("/health", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"healthy":   false,
				"dbHealthy": false,
				"error":     err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"healthy":   true,
			"dbHealthy": true,
		})
	})

	messages.Init(engine, client.Database(envVars.MongoDB.DatabaseName))

	engine.Run()
}

func connectToDatabase(envVars *Environment) (*mongo.Client, error) {
	connectCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(connectCtx, options.Client().ApplyURI(envVars.MongoDB.ConnectionURI).SetAuth(options.Credential{
		AuthMechanism: envVars.MongoDB.AuthMechanism,
		AuthSource:    envVars.MongoDB.AuthDatabase,
		Username:      envVars.MongoDB.Username,
		Password:      envVars.MongoDB.Password,
	}))

	if err != nil {
		return nil, err
	}

	pingCtx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}
