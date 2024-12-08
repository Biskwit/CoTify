package utils

import (
	"bytes"
	"encoding/json"
	"os"
	"strconv"

	"io/ioutil"
	"net/http"

	"github.com/Biskwit/CoTify/types"
	log "github.com/sirupsen/logrus"
)

func MakeCoTIteration(url string, api_key string, model string, query string) (*string, error) {
	requestBody := types.ChatCompletionsRequest{
		Model:       model,
		MaxTokens:   300,
		Temperature: 0.2,
		ResponseFormat: types.ResponseFormat{
			Type: "json_object",
		},
		Messages: []types.Message{
			{
				Role: "system",
				Content: `You are an expert AI assistant that explains your reasoning step by step. For each step, provide a title that describes what you're doing in that step, along with the content. Decide if you need another step or if you're ready to give the final answer. Respond in JSON format with 'title', 'content', and 'next_action' (either 'continue' or 'final_answer') keys. USE AS MANY REASONING STEPS AS POSSIBLE. AT LEAST 3. BE AWARE OF YOUR LIMITATIONS AS AN LLM AND WHAT YOU CAN AND CANNOT DO. IN YOUR REASONING, INCLUDE EXPLORATION OF ALTERNATIVE ANSWERS. CONSIDER YOU MAY BE WRONG, AND IF YOU ARE WRONG IN YOUR REASONING, WHERE IT WOULD BE. FULLY TEST ALL OTHER POSSIBILITIES. YOU CAN BE WRONG. WHEN YOU SAY YOU ARE RE-EXAMINING, ACTUALLY RE-EXAMINE, AND USE ANOTHER APPROACH TO DO SO. DO NOT JUST SAY YOU ARE RE-EXAMINING. USE AT LEAST 3 METHODS TO DERIVE THE ANSWER. USE BEST PRACTICES.
						  Example of a valid JSON response:
						  {
						  	"title": "Identifying Key Information",
							"content": "To begin solving this problem, we need to carefully examine the given information and identify the crucial elements that will guide our solution process. This involves...",
							"next_action": "continue"
						}`,
			},
			{
				Role:    "user",
				Content: query,
			},
			{
				Role:    "assistant",
				Content: "Thank you! I will now think step by step following my instructions, starting at the beginning after decomposing the problem.",
			},
		},
	}
	iterationsStr := os.Getenv("COT_ITERATIONS")
	iterations, err := strconv.Atoi(iterationsStr)
	if err != nil {
		log.Error("Error converting COT_ITERATIONS to int:", err)
		return nil, err
	}

	for i := 0; i < iterations; i++ {

		// Marshal the payload to JSON
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			log.Error("Error marshalling JSON:", err)
			return nil, err
		}

		// Create a new HTTP POST request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Error("Error creating HTTP request:", err)
			return nil, err
		}

		// Set the necessary headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+api_key)

		// Use the default HTTP client to send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Error("Error sending HTTP request:", err)
			return nil, err
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error("Error reading response body:", err)
			return nil, err
		}

		var chatResponse types.ChatCompletionResponse
		err = json.Unmarshal(body, &chatResponse)
		if err != nil {
			log.Error("Error unmarshalling response JSON:", err)
			return nil, err
		}

		requestBody.Messages = append(requestBody.Messages, types.Message{
			Role:    "assistant",
			Content: chatResponse.Choices[0].Message.Content,
		})
	}

	requestBody.Messages = append(requestBody.Messages, types.Message{
		Role:    "user",
		Content: "Please provide the final answer based solely on your reasoning above. Do not use JSON formatting. Only provide the text response without any titles or preambles. Retain any formatting as instructed by the original prompt, such as exact formatting for free response or multiple choice.",
	})
	requestBody.ResponseFormat.Type = "text"

	// Marshal the payload to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Error("Error marshalling JSON:", err)
		return nil, err
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error("Error creating HTTP request:", err)
		return nil, err
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api_key)

	// Use the default HTTP client to send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error sending HTTP request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading response body:", err)
		return nil, err
	}

	var chatResponse types.ChatCompletionResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		log.Error("Error unmarshalling response JSON:", err)
		return nil, err
	}

	return &chatResponse.Choices[0].Message.Content, nil
}
