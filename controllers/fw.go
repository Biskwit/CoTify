package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Forward(c *gin.Context) {
	log.Info("Forwarding request to OpenAI")

	if c.Request.URL.Path == "/chat/completions" {
		ChatCompletions(c)
		return
	}

	req, err := http.NewRequest(c.Request.Method, os.Getenv("OPENAI_API_BASE")+c.Request.URL.Path, c.Request.Body)
	if err != nil {
		log.Error("Error creating new request: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error making request to OpenAI: ", err)
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "Failed to connect to OpenAI API"})
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	contentType := resp.Header.Get("Content-Type")
	c.DataFromReader(resp.StatusCode, resp.ContentLength, contentType, resp.Body, nil)
}
