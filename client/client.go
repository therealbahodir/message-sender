package client

import (
	"log"
	"google.golang.org/grpc"
	"github.com/therealbahodir/message-sender/proto"
	"context"
	"fmt"
	"os"
    "github.com/joho/godotenv"
)

func Stub (message, priority string) (status int){
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Problem with loading env: %v", err)
    }

    port := os.Getenv("GRPC_SERVER_PORT")
	conn, err := grpc.Dial(fmt.Sprintf("localhost%s", port), grpc.WithInsecure())
	if err != nil {
		log.Printf("Problem with connecting client to server: %s", err)
		return 400
	}

	defer conn.Close()

	stub := proto.NewBotClient(conn)

	_, err = stub.Sender(
		context.Background(), 
		&proto.Content{Text:message,Priority: priority},
	)

		

	if err != nil {
		log.Printf("Problem with sending request to server: %v", err)
		return 400
	}
	return 200
}