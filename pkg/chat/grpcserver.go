package chat

import (
	"context"
	"github.com/google/uuid"
	"grpc3/pkg/api"
	"grpc3/pkg/database"
	"log"
)

type GRPCserver struct {
	api.UnimplementedChatServer
	DB database.Database
}

func (s *GRPCserver) AddMessage(ctx context.Context, req *api.Message) (*api.MessageResponce, error) {
	id := req.GetId()
	message := req.GetMessage()
	date := req.GetTime()
	resp, err := s.DB.AddMessage(ctx, req)
	if err != nil {
		log.Println("Error entering values in DB", err)
	}
	log.Println("Response:", resp)
	return &api.MessageResponce{
		Id:        id,
		Message:   message,
		TimeSaved: date,
	}, nil
}

func (s *GRPCserver) AddChat(ctx context.Context, req *api.ChatId) (*api.ChatId, error) {
	id := uuid.New().String()
	req = &api.ChatId{Id: id}
	save, err := s.DB.AddChat(ctx, req)
	if err != nil {
		log.Println("Error saving chat in DB:", err)
	}
	log.Println(save)
	return &api.ChatId{Id: id}, nil
}

func (s *GRPCserver) GetChat(ctx context.Context, in *api.ChatId) (*api.GetChatResponse, error) {
	resp, err := s.DB.GetChat(ctx, in)
	if err != nil {
		log.Println("error getting chat", err)
		return nil, err
	}

	// Создаем массив сообщений из ответа БД
	var messages []*api.Message
	for _, msg := range resp.Messages {
		// Преобразуем время из Timestamp protobuf во время Go
		//timeSaved, err := msg.Time.AsTime()
		//if err != nil {
		//	log.Println("Error converting timestamp:", err)
		//	continue
		//}
		// Создаем новый объект сообщения с преобразованным временем
		newMsg := &api.Message{
			Id: msg.Id,
			//Name:    msg.Name,
			//Message: msg.Message,
			//Time:    timestamppb.New(timeSaved),
		}
		messages = append(messages, newMsg)
	}

	// Возвращаем ответ с массивом сообщений
	return &api.GetChatResponse{Messages: messages}, nil
}

//func (s *GRPCserver) StreamChat(req *api.ChatId, stream api.Chat_StreamChatServer) error {
//	idChat := req.GetId()
//	log.Print("receive get-chat request with chat id: ", idChat)
//
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//	res, err := s.DB.StreamChat(ctx, idChat, stream)
//	if err != nil {
//		log.Println("cannot find messages in DB: ", err)
//		return err
//	}
//	err = stream.Send(res)
//	if err != nil {
//		log.Printf("cannot send messages from chat id: %s", idChat)
//		return err
//	}
//	log.Println("sent message from chat: ", idChat)
//	return nil
//}
