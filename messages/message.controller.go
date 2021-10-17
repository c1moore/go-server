package messages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type messageController struct {
	service *messageService
}

func newMessageController(service *messageService) *messageController {
	return &messageController{
		service: service,
	}
}

func (this *messageController) GetMessages(context *gin.Context) {
	context.JSON(http.StatusOK, this.service.LoadMessages())
}

func (this *messageController) PostMessage(context *gin.Context) {
	var message *Message

	if err := context.BindJSON(&message); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, this.service.AddMessage(*message))
}
