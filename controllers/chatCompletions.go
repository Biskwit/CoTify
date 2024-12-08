package controllers

import (
	"github.com/Biskwit/CoTify/services"
	"github.com/Biskwit/CoTify/types"
	"github.com/gin-gonic/gin"
)

func ChatCompletions(c *gin.Context) {
	var body types.ChatCompletionsRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatus(406)
		return
	}
	services.ChatCompletions(c, body)
}
