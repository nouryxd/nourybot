package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	// Bot
	Username  string
	BotUserId string
	Oauth     string
	// Admin
	AdminUsername string
	AdminUserId   string

	// DB
	MySQLURI string
	MongoURI string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	cfg := &Config{
		// Bot
		Username:  os.Getenv("BOT_USERNAME"),
		Oauth:     os.Getenv("BOT_PASS"),
		BotUserId: os.Getenv("BOT_USER_ID"),

		// Admin
		AdminUsername: os.Getenv("ADMIN_USER_NAME"),
		AdminUserId:   os.Getenv("ADMIN_USER_ID"),

		// DB
		MySQLURI: os.Getenv("MYSQL_URI"),
		MongoURI: os.Getenv("MONGO_URI"),
	}

	log.Info("Config loaded succesfully")

	return cfg
}
