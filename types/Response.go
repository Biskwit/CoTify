package types

type Output struct {
	Inputs  Inputs    `json:"inputs"`
	Outputs []Outputs `json:"outputs"`
}

type Inputs struct {
	InputValue string `json:"input_value"`
}

type Outputs struct {
	Results   Results           `json:"results"`
	Artifacts Artifacts         `json:"artifacts"`
	Outputs   Outputs2          `json:"outputs"`
	Logs      Logs              `json:"logs"`
	Messages  []MessageResponse `json:"messages,omitempty"`
}

type Results struct {
	Message MessageResponse `json:"message"`
}

type Artifacts struct {
	Message    string   `json:"message"`
	Sender     string   `json:"sender"`
	SenderName string   `json:"sender_name"`
	Files      []string `json:"files"`
	Type       string   `json:"type"`
	StreamUrl  string   `json:"stream_url"`
}

type Outputs2 struct {
	Message NestedMessage `json:"message"`
}

type NestedMessage struct {
	Message MessageResponse `json:"message"`
	Type    string          `json:"type"`
}

type Logs struct {
	Message []string `json:"message"`
}

type MessageResponse struct {
	Data       Data     `json:"data"`
	Sender     string   `json:"sender"`
	SenderName string   `json:"sender_name"`
	SessionID  string   `json:"session_id"`
	Files      []string `json:"files"`
	Timestamp  string   `json:"timestamp"`
	FlowID     string   `json:"flow_id"`
}

type Data struct {
	Text string `json:"text"`
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int           `json:"index"`
	Message      OpenAIMessage `json:"message"`
	Logprobs     *int          `json:"logprobs"` // use *int to handle null values
	FinishReason string        `json:"finish_reason"`
}

type CompletionTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}

type Usage struct {
	PromptTokens            int                     `json:"prompt_tokens"`
	CompletionTokens        int                     `json:"completion_tokens"`
	TotalTokens             int                     `json:"total_tokens"`
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`
}

type ChatCompletionResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	Error             string   `json:"error,omitempty"`
}
