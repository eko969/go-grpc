package main

import (
	"log"
	"net"

	"context"

	"github.com/eko9x9/go-grpc/chat/chatpb"
	"google.golang.org/grpc"
)

type Server struct {
	chatpb.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *chatpb.Message) (*chatpb.Message, error) {
	return &chatpb.Message{Body: "Hello from chatpb"}, nil
}

func main() {
	list, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen port 9000 %v", err)
	}

	grpcServer := grpc.NewServer()

	chatServer := Server{}
	chatpb.RegisterChatServiceServer(grpcServer, &chatServer)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Grpc failed to listen on port 9000: %v", err)
	}

}
