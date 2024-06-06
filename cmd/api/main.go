package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"grpc3/pkg/api"
	"grpc3/pkg/chat"
	"grpc3/pkg/database"
	"log"
	"net"
	"time"
)

func main() {
	client := GetMongoClient()
	db := database.NewMongoClient(client, "chat")

	s := grpc.NewServer()
	srv := &chat.GRPCserver{DB: db}
	api.RegisterChatServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error running server:", err)
	}
	if err = s.Serve(l); err != nil {
		log.Fatal("Error: ", err)
	}
}

func GetMongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return client
}
