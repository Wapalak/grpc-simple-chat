package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	api2 "grpc3/pkg/api"
	"log"
	"os"
)

func main() {
	args := os.Args[1:] // Получаем аргументы командной строки, начиная с первого (индекс 0 это имя исполняемого файла)
	if len(args) != 1 {
		fmt.Println("Использование: go run main.go 1")
		return
	}
	id := args[0]
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error dialing up:", err)
	}
	c := api2.NewChatClient(conn)
	ctx := context.Background()
	chatId := &api2.ChatId{Id: id}
	res, err := c.AddChat(ctx, chatId)
	if err != nil {
		log.Fatal("Error", err)
	}
	log.Println("Response from server", res)
}
