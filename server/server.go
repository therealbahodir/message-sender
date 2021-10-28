package main

import (
	"net"
    "google.golang.org/grpc"
    "github.com/therealbahodir/message-sender/proto"
    "context"
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Server struct {
    proto.UnimplementedBotServer
}

func (*Server) Sender(ctx context.Context,input *proto.Content) (*proto.Empty, error) {
    
    return &proto.Empty{}, nil
    
}


func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Problem with loading env: %v", err)
    }

    port := os.Getenv("HTTP_PORT")

	lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Problem with connecting to tcp: %v", err)
    }

    s := grpc.NewServer()
    proto.RegisterBotServer(s, &Server{})
    log.Println("Server is running in port :6060")
    err = s.Serve(lis)
    if err != nil {
        log.Fatalf("Problem with connecting to GRPC Server: %v", err)
    }
}