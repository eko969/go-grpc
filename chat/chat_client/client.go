package main

import (
	"context"
	"log"

	"github.com/eko9x9/go-grpc/chat/chatpb"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could't connect server %v", err)
	}

	defer conn.Close()

	c := chatpb.NewChatServiceClient(conn)

	message := chatpb.Message{
		Body: "Hello from the client",
	}
	resp, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling SayHello %v", err)
	}

	log.Printf("Respon from the server %v", resp)

}
