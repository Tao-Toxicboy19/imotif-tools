package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Model     string
	Provider  string
	APIKey    string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	return Config{
		Model:    os.Getenv("AI_MODEL"),
		Provider: os.Getenv("AI_PROVIDER"),
		APIKey:   os.Getenv("GOOGLE_API_KEY"),
	}
}
