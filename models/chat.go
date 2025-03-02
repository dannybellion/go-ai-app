package models

import (
	"log"

	"github.com/dannybellion/go-ai-app/database"
)

// Chat struct (represents a row in the chat_history table)
type Chat struct {
	ID       int    `json:"id"`
	Prompt   string `json:"prompt"`
	Response string `json:"response"`
}

// Save chat to database
func SaveChat(prompt, response string) {
	_, err := database.DB.Exec("INSERT INTO chat_history (prompt, response) VALUES (?, ?)", prompt, response)
	if err != nil {
		log.Println("Error saving chat:", err)
	}
}

// Retrieve all chat history
func GetChatHistory() ([]Chat, error) {
	rows, err := database.DB.Query("SELECT id, prompt, response FROM chat_history ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []Chat
	for rows.Next() {
		var chat Chat
		rows.Scan(&chat.ID, &chat.Prompt, &chat.Response)
		history = append(history, chat)
	}
	return history, nil
}