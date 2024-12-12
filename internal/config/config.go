package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI       string
	ServerPort     string
	DatabaseName   string
	CollectionName string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		MongoURI:       getEnv("MONGO_URI", "mongodb://localhost:27017"),
		ServerPort:     getEnv("SERVER_PORT", "3000"),
		DatabaseName:   getEnv("DATABASE_NAME", "service_registry"),
		CollectionName: getEnv("COLLECTION_NAME", "services"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
