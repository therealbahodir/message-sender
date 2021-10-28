package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
)


type Message struct {
	Text string   			`json:"text"`
	Priority string 		`json:"priority"`
}

// @Summary MessageSender
// @Description send message
// @ID message_sender
// @Accept json
// @Produce json
// @Param input body Message true "message_content"
// @Success 200
// @Failure 400
// @Router /send [post]
func SendMessageAPI(ctx *gin.Context) {

	var newMessage Message
	err := ctx.BindJSON(&newMessage)
	if err != nil {
		log.Fatalf("Problem with getting message from API Server: %v", err)
	}

   	log.Println(newMessage)

    ctx.AbortWithStatus(200)

}
