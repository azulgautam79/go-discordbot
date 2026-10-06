package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	Env       string
	Token     string
	BotPrefix string
}

func MustLoad() Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//! Port
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is required")
	}

	//! Env
	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required")
	}

	//! Token
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("TOKEN is required")
	}

	//! BotPrefix
	botPrefix := os.Getenv("BOTPREFIX")
	if botPrefix == "" {
		panic("BOTPREFIX is required")
	}

	return Config{
		Port:      port,
		Env:       env,
		Token:     token,
		BotPrefix: botPrefix,
	}
}
