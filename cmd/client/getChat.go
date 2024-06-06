package main

//import (
//	"context"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"google.golang.org/protobuf/types/known/timestamppb"
//	"grpc3/pkg/api"
//	"io"
//)
//
//type chatStreamClient struct {
//	cursor mongo.Cursor
//}
//
//// Метод Recv для получения сообщения из курсора
//func (c *chatStreamClient) Recv() (*api.Message, error) {
//	if !c.cursor.Next(context.Background()) {
//		return nil, io.EOF
//	}
//
//	var result bson.M
//	if err := c.cursor.Decode(&result); err != nil {
//		return nil, err
//	}
//
//	// Создаем сообщение из результата запроса
//	// Помните, что этот код нужно адаптировать под вашу структуру данных в MongoDB
//	message := &api.Message{
//		Id:      result["_id"].(string),
//		Name:    result["name"].(string),
//		Message: result["message"].(string),
//		Time:    result["time"].(timestamppb.Timestamp), // Используйте timestamppb.Timestamp
//	}
//
//	return message, nil
//}
