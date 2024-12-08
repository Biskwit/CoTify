package types

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

type ChatCompletionsRequest struct {
	Model          string         `json:"model"`
	Messages       []Message      `json:"messages"`
	Stream         bool           `json:"stream"`
	MaxTokens      int            `json:"max_tokens"`
	Temperature    float64        `json:"temperature"`
	ResponseFormat ResponseFormat `json:"response_format"`
}
