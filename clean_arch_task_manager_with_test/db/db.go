package db

import (
	"cleantaskmanager/mongo"
	"context"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
)

var Client mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		// Handle error (optional)
		panic("Error loading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.NewClient(os.Getenv("MONGO_URI"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
}
