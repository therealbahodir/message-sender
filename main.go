package main 

import (
	"os"
    "github.com/joho/godotenv"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/therealbahodir/message-sender/handlers"
    _"github.com/therealbahodir/message-sender/docs"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

// @title Message Sender Bot
// @version 1.0
// @description Telegram Bot which sends messages to channels and groups
// @host localhost:8000
// @BasePath /
func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Problem with loading env: %v", err)
    }

    port := os.Getenv("HTTP_PORT")


    router := gin.Default()

    router.POST("/send", handlers.SendMessageAPI)
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


    log.Println("Listening and serving HTTP on localhost:", port)
    err = router.Run(port)
    if err != nil {
        log.Fatalf("Problem with connecting gateway to port: %v",err)
    }
}