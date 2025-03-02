package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	serverURL = "http://localhost:8080"
)

func main() {
	fmt.Println("AI Chat CLI")
	fmt.Println("Type 'exit' or 'quit' to exit")
	fmt.Println("Type 'history' to view chat history")
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if input == "exit" || input == "quit" {
			break
		}
		
		if input == "history" {
			showHistory()
			continue
		}
		
		if input == "" {
			continue
		}
		
		response, err := sendChat(input)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		
		fmt.Printf("\nAI: %s\n", response)
	}
}

func sendChat(prompt string) (string, error) {
	requestBody, err := json.Marshal(map[string]string{
		"prompt": prompt,
	})
	if err != nil {
		return "", err
	}
	
	resp, err := http.Post(serverURL+"/chat", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("server connection error: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("server error: %s - %s", resp.Status, string(bodyBytes))
	}
	
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	
	response, ok := result["response"].(string)
	if !ok {
		return "", fmt.Errorf("unexpected response format")
	}
	
	return response, nil
}

func showHistory() {
	resp, err := http.Get(serverURL + "/history")
	if err != nil {
		fmt.Printf("Error fetching history: %s\n", err)
		return
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: server returned %s\n", resp.Status)
		return
	}
	
	var history []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&history); err != nil {
		fmt.Printf("Error parsing history: %s\n", err)
		return
	}
	
	fmt.Println("\n=== Chat History ===")
	for i, entry := range history {
		fmt.Printf("\n--- Conversation %d ---\n", i+1)
		fmt.Printf("User: %s\n", entry["prompt"])
		fmt.Printf("AI: %s\n", entry["response"])
	}
}