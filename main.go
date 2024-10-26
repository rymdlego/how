package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIRequest struct {
	Model     string        `json:"model"`
	Messages  []ChatMessage `json:"messages"`
	MaxTokens int           `json:"max_tokens"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a question or input text.")
		return
	}

	// Combine all arguments into a single question string with additional instructions
	question := strings.Join(os.Args[1:], " ") + " Answer with only the command-line content, no extra text."

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("API key missing. Please set the OPENAI_API_KEY environment variable.")
		return
	}

	// Prepare the chat-based request payload
	requestBody, err := json.Marshal(OpenAIRequest{
		Model: "gpt-3.5-turbo",
		Messages: []ChatMessage{
			{Role: "user", Content: question},
		},
		MaxTokens: 50,
	})
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
		return
	}

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var openAIResponse OpenAIResponse
	if err := json.Unmarshal(body, &openAIResponse); err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	// Output the response
	if len(openAIResponse.Choices) > 0 {
		fmt.Println(openAIResponse.Choices[0].Message.Content)
	} else {
		fmt.Println("No response received.")
	}
}
