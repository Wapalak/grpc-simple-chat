package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"grpc3/pkg/api"
	"log"
	"sync"
)

type MongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
	mutex      sync.RWMutex
	data       map[string]*api.Message
}

func (m *MongoDB) StreamChat(ctx context.Context, chatID string, stream api.Chat_StreamChatServer) (*api.Message, error) {
	//TODO implement me
	panic("implement me")
}

func NewMongoClient(client *mongo.Client, collectionName string) *MongoDB {
	db := client.Database("grpctest")
	collection := db.Collection(collectionName)
	return &MongoDB{
		client:     client,
		collection: collection,
	}
}

func (m *MongoDB) AddMessage(ctx context.Context, req *api.Message) (*api.MessageResponce, error) {
	filter := bson.M{"id": req.GetId()}
	messageDoc := bson.M{
		"name":    req.GetName(),
		"message": req.GetMessage(),
		"time":    req.GetTime(),
	}
	update := bson.M{"$push": bson.M{"messages": messageDoc}}

	if _, err := m.collection.UpdateOne(ctx,
		filter,
		update,
		options.Update().SetUpsert(true),
	); err != nil {
		log.Fatal("Error inserting data:", err)
	}

	return &api.MessageResponce{}, nil
}

func (m *MongoDB) AddChat(ctx context.Context, in *api.ChatId, opts ...grpc.CallOption) (*api.ChatId, error) {
	if _, err := m.collection.InsertOne(ctx, in); err != nil {
		log.Println("Error creating chat:", err)
		return nil, err
	}
	return in, nil
}

func (m *MongoDB) GetChat(ctx context.Context, in *api.ChatId, opts ...grpc.CallOption) (*api.GetChatResponse, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// Читаем сообщения из базы данных
	messagesFromDB, err := m.getMessagesFromChat(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	if len(messagesFromDB) == 0 {
		log.Println("no messages in chat: ", in.GetId())
		return &api.GetChatResponse{
			Messages: []*api.Message{},
		}, nil
	}
	log.Println("got messages from chat: ", in.GetId())

	// Преобразуем каждое сообщение в формат Message
	var messages []*api.Message
	for _, msgDB := range messagesFromDB {
		msg := &api.Message{
			Id:      msgDB.Id,
			Name:    msgDB.Name,
			Message: msgDB.Message,
			Time:    msgDB.Time,
		}

		// Преобразуем время в формат Timestamp
		//timeProto, err := ptypes.TimestampProto(msgDB.Time)
		//if err != nil {
		//	return nil, err
		//}
		//msg.Time = timeProto

		// Добавляем сообщение в массив messages
		messages = append(messages, msg)
	}

	// Создаем объект GetChatResponse с заполненным массивом messages и возвращаем его
	return &api.GetChatResponse{
		Messages: messages,
	}, nil
}

func (m *MongoDB) getMessagesFromChat(ctx context.Context, chatID string) ([]api.Message, error) {
	// Создаем фильтр для поиска сообщений по chat_id
	filter := bson.M{"id": chatID, "messages": bson.M{"$exists": true, "$ne": bson.A{}}}

	// Выполняем поиск сообщений в коллекции MongoDB
	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		log.Println("Error finding messages:", err)
		return nil, err
	}
	log.Println("found chat:", chatID)
	log.Println(cursor.Current)
	defer cursor.Close(ctx)

	// Создаем слайс для хранения найденных сообщений
	var messages []api.Message

	// Итерируемся по результатам поиска
	for cursor.Next(ctx) {
		log.Println("starting adding messages")
		var msg api.Message
		// Декодируем текущее сообщение в структуру Message
		if err := cursor.Decode(&msg); err != nil {
			log.Println("Error decoding message:", err)
			continue
		}
		// Добавляем сообщение в слайс
		messages = append(messages, msg)
		log.Println("added message to slice: ", messages)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}
	log.Println("messages: ", messages)
	return messages, nil
}

//func (m *MongoDB) StreamChat(ctx context.Context, chatID string, stream api.Chat_StreamChatServer) (*api.Message, error) {
//	m.mutex.Lock()
//	defer m.mutex.Unlock()
//
//	for _, message := range m.data {
//		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
//			log.Print("context is cancelled")
//			return nil, fmt.Errorf("context is canceled")
//		}
//
//		filter := bson.M{"id": chatID, "messages": bson.M{"$exists": true, "$ne": bson.A{}}}
//
//		// Выполняем поиск сообщений в коллекции MongoDB
//		cursor, err := m.collection.Find(ctx, filter)
//		if err != nil {
//			log.Println("Error finding messages:", err)
//			return nil, err
//		}
//		log.Println("found chat:", chatID)
//		log.Println(cursor.Current)
//		defer cursor.Close(ctx)
//
//	}
//}
