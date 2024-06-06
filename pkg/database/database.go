package database

import (
	"context"
	"google.golang.org/grpc"
	"grpc3/pkg/api"
)

type Database interface {
	AddMessage(ctx context.Context, req *api.Message) (*api.MessageResponce, error)
	AddChat(ctx context.Context, in *api.ChatId, opts ...grpc.CallOption) (*api.ChatId, error)
	GetChat(ctx context.Context, in *api.ChatId, opts ...grpc.CallOption) (*api.GetChatResponse, error) // Другие методы работы с базой данных
	getMessagesFromChat(ctx context.Context, chatID string) ([]api.Message, error)

	//StreamChat(ctx context.Context, chatID string, stream api.Chat_StreamChatServer) (*api.Message, error)
}
