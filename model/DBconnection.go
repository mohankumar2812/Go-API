package model

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

func DBConnect() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	uri := os.Getenv("mongoDB")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second * 20);
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database("student")
	fmt.Println("db connect successfully");

}
