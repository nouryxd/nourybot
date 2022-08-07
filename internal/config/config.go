package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TwitchUsername string
	TwitchOauth    string
	Environment    string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	twitchUsername := os.Getenv("TWITCH_USERNAME")
	twitchOauth := os.Getenv("TWITCH_OAUTH")
	environment := "Development"

	return &Config{twitchUsername, twitchOauth, environment}
}
