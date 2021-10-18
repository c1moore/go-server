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

func (ctl *messageController) GetMessages(context *gin.Context) {
	messages, err := ctl.service.LoadMessages()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, messages)
}

func (ctl *messageController) PostMessage(context *gin.Context) {
	var message *Message

	if err := context.BindJSON(&message); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := ctl.service.AddMessage(message); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, message)
}
