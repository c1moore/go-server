package messages

import "github.com/gin-gonic/gin"

func Init(engine *gin.Engine) {
	repo := newMessageRepo()
	service := newMessageService(repo)
	ctl := newMessageController(service)

	engine.GET("/messages", ctl.GetMessages)

	engine.POST("/messages", ctl.PostMessage)
}
