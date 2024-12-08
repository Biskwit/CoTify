package services

import (
	"net/http"
	"os"

	"github.com/Biskwit/CoTify/types"
	"github.com/Biskwit/CoTify/utils"
	u "github.com/Biskwit/CoTify/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ChatCompletions(c *gin.Context, request types.ChatCompletionsRequest) {

	apiKey := c.GetString("apiKey")

	res, err := u.MakeCoTIteration(os.Getenv("OPENAI_API_BASE")+c.Request.URL.Path, apiKey, request.Model, request.Messages[0].Content)
	if err != nil {
		log.Error("Error making CoT iteration: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	var response types.ChatCompletionResponse = types.ChatCompletionResponse{}
	response.Usage.PromptTokens = utils.Tokenize(request.Messages[len(request.Messages)-1].Content)
	response.Usage.CompletionTokens = utils.Tokenize(*res)
	response.Choices = append(response.Choices, types.Choice{
		Index: 0,
		Message: types.OpenAIMessage{
			Role:    "assistant",
			Content: *res,
		},
	})
	c.JSON(http.StatusOK, response)
}
