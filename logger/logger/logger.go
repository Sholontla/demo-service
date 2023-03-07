package logger

import (
	"context"
	"log"
	"time"

	"service/logger/service/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Config struct {
	LoggerModel LogModel
}

func MongoConn() *Config {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient
	ctxBack := context.Background()
	ctx, cancel := context.WithTimeout(ctxBack, 15*time.Second)
	defer cancel()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		LoggerModel: New(client),
	}

	return &app
}

func connectToMongo() (*mongo.Client, error) {

	var (
		mongoUrl, user, pass = config.DBConfig()
		clientOptions        = options.Client().ApplyURI(mongoUrl)
		ctx                  = context.TODO()
	)
	clientOptions.SetAuth(options.Credential{
		Username: user,
		Password: pass,
	})

	conn, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Printf("Error connectexit to MongoDB: %s", err)
		return nil, err
	}

	return conn, nil
}
