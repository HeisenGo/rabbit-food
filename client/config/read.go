package config

import (
	"client/services"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnvVars(filePath string) {
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func readAPIService() services.Service {
	host := os.Getenv("API_SERVICE_HOST")
	port := os.Getenv("API_SERVICE_PORT")
	service := services.GetAPIService(host, port)
	return service
}

func ReadConfig(envFilePath string) IConfig {
	loadEnvVars(envFilePath)
	service := readAPIService()
	config := NewConfig(service)
	return config
}
