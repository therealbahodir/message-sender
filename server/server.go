package main

import (
	"net"
    "google.golang.org/grpc"
    "github.com/therealbahodir/message-sender/proto"
    "github.com/therealbahodir/message-sender/bot"
    "github.com/therealbahodir/message-sender/handlers"
    "context"
    "log"
    "os"
    "github.com/joho/godotenv"
    "time"
)

var List []handlers.Message


type Server struct {
    proto.UnimplementedBotServer
}

func (*Server) Sender(ctx context.Context,input *proto.Content) (*proto.Empty, error) {
   
    var message handlers.Message

	message.Priority = input.GetPriority()
	message.Text = input.GetText()

	List = append(List, message)

    return &proto.Empty{}, nil
}


func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Problem with loading env: %v", err)
    }

    port := os.Getenv("GRPC_SERVER_PORT")


    go CuncurrentSending()

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


func CuncurrentSending() {
	var isSend bool
    for {
        isSend = false
        for i, message := range List{
            if message.Priority == "high" {
                err := bot.BotService(List[i].Text)
                if err != nil {
                    log.Fatalf("Problem with sending message to bot: %v",err) 
                } else {
                    isSend = true
                    List = Remove(List, i)
                    time.Sleep(time.Second * 40) 

                    break
                } 
            }
        }
        
        if isSend {
            continue
        }

        for i, message := range List{
            if message.Priority == "medium" {
                err := bot.BotService(List[i].Text)
                if err != nil {
                    log.Fatalf("Problem with sending message to bot: %v",err) 
                } else {
                    isSend = true
                    List = Remove(List, i)
                    time.Sleep(time.Second * 40) 

                    break
                } 
            }
        }
        if isSend {
            continue
        }
        for i, message := range List{
        	if message.Priority == "low" {
                err := bot.BotService(List[i].Text)
                if err != nil {
                    log.Fatalf("Problem with sending message to bot: %v",err) 
                } else {
                    isSend = true
                    List = Remove(List, i)
                    time.Sleep(time.Second * 40) 

                    break
                } 
            }
        }


    }
}

func Remove(s []handlers.Message, i int) []handlers.Message {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}