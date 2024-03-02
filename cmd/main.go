package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "go-chat/pkg/chat/chat_v1"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const serverPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	log.Printf("Creating chat")

	return &desc.CreateChatResponse{
		Id: 1,
	}, nil
}

func (s *server) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*empty.Empty, error) {
	log.Printf("Delete chat")

	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	log.Printf("Send message")

	return &emptypb.Empty{}, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterChatV1Server(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
