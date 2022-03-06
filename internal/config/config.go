package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Username  string
	Oauth     string
	BotUserId string
	MongoURI  string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	cfg := &Config{
		Username:  os.Getenv("BOT_USERNAME"),
		Oauth:     os.Getenv("BOT_PASS"),
		BotUserId: os.Getenv("BOT_USER_ID"),
		MongoURI:  os.Getenv("MONGO_URI"),
	}

	log.Info("Config loaded succesfully")

	return cfg
}
