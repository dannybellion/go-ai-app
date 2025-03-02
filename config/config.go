package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	OpenAIKey string
	OpenAIURL = "https://api.openai.com/v1/chat/completions"
	OpenAIModel = "gpt-4o-mini"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env variables")
	}
	OpenAIKey = os.Getenv("OPENAI_API_KEY")
}