package main

// https://www.youtube.com/watch?v=XpQetNNNqvo&list=PLVDJsRQrTUz7-fSzZWtF726st5AYQc57A&index=1&t=511s phut thu 22:43
import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// fmt.Println("hello world")
	Connect()
}
func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://host.docker.internal:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection ok")

}
