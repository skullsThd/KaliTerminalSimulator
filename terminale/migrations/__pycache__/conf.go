package main

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Port         string
	DatabaseURL  string
	APISecretKey string
	Timeout      time.Duration
}

func LoadConfig() *Config {
	cfg := &Config{
		Port:         getEnv("PORT", "8000"),
		DatabaseURL:  getEnv("DATABASE_URL", ""),
		APISecretKey: getEnv("API_SECRET_KEY", ""),
	}

	timeout, err := time.ParseDuration(getEnv("TIMEOUT", "10s"))
	if err != nil {
		log.Fatal("Invalid timeout duration")
	}
	cfg.Timeout = timeout

	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {
	// Carica la configurazione
	config := LoadConfig()

	// Utilizza la configurazione nel tuo programma
	log.Println("Port:", config.Port)
	log.Println("Database URL:", config.DatabaseURL)
	log.Println("API Secret Key:", config.APISecretKey)
	log.Println("Timeout:", config.Timeout)
}
