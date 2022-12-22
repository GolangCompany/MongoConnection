package main

//
import (
	"context"
	"fmt"
	"log"
	"time"

	//  "os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://<Username>:<Password>@cluster0.0aeeqhe.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failes to connect mongodb")
		fmt.Println(err)
	}
	if err == nil {
		fmt.Println("Successfully Connected to the mongodb")
	}
}

//
