package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dannybellion/go-ai-app/config"
)

// HTTPClient provides a reusable client for making HTTP requests
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates a new HTTP client with sensible defaults
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// OpenAIRequest represents the structure of an OpenAI API request
type OpenAIRequest struct {
	Model    string                  `json:"model"`
	Messages []OpenAIRequestMessage  `json:"messages"`
}

// OpenAIRequestMessage represents a message in an OpenAI conversation
type OpenAIRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse represents the structure of an OpenAI API response
type OpenAIResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// CallOpenAI sends a request to the OpenAI API and returns the response
func (c *HTTPClient) CallOpenAI(prompt string) (string, error) {
	reqBody := OpenAIRequest{
		Model: config.OpenAIModel,
		Messages: []OpenAIRequestMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}
	
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}
	
	req, err := http.NewRequest("POST", config.OpenAIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.OpenAIKey)
	
	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API error: %s - %s", resp.Status, string(bodyBytes))
	}
	
	var openAIResp OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}
	
	if len(openAIResp.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}
	
	return openAIResp.Choices[0].Message.Content, nil
}
