package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	api2 "grpc3/pkg/api"
	"log"
	"os"
	"time"
)

func main() {
	args := os.Args[1:] // Получаем аргументы командной строки, начиная с первого (индекс 0 это имя исполняемого файла)
	if len(args) != 3 {
		fmt.Println("Использование: go run main.go uuid name message")
		return
	}
	id := args[0]
	name := args[1]
	message := args[2]
	now := time.Now() // Получаем текущее время

	// Создаем timestamp.Timestamp с помощью отформатированного времени
	date := &timestamp.Timestamp{
		Seconds: now.Unix(),
		Nanos:   int32(now.Nanosecond()),
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error establishing connection: ", err)
	}

	c := api2.NewChatClient(conn)

	res, err := c.AddMessage(context.Background(), &api2.Message{
		Id:      id,
		Name:    name,
		Message: message,
		Time:    date,
	})
	if err != nil {
		log.Fatal("error: ", err)
	}
	log.Println(res.GetMessage())
}
