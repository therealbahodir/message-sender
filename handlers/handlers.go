package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
)


type Message struct {
	Text string   			`json:"text"`
	Priority string 		`json:"priority"`
}


func SendMessageAPI(ctx *gin.Context) {

	var newMessage Message
	err := ctx.BindJSON(&newMessage)
	if err != nil {
		log.Fatalf("Problem with getting message from API Server: %v", err)
	}

   	log.Println(newMessage)

    ctx.AbortWithStatus(200)

}
