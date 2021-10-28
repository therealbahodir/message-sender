package main 

import (
	"os"
    "github.com/joho/godotenv"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/therealbahodir/message-sender/handlers"
)


func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Problem with loading env: %v", err)
    }

    port := os.Getenv("HTTP_PORT")


    router := gin.Default()

    router.POST("/send", handlers.SendMessageAPI)


    log.Println("Listening and serving HTTP on localhost:", port)
    err = router.Run(port)
    if err != nil {
        log.Fatalf("Problem with connecting gateway to port: %v",err)
    }
}