package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetMongoDBCredentials() (string, error) {
	err := godotenv.Load()

	if err != nil {
		return "", err
	}

	mongoURL := os.Getenv("MONGODB_URL")
	mongoPort := os.Getenv("MONGODB_PORT")
	mongoUser := os.Getenv("MONGODB_USERNAME")
	mongoPass := os.Getenv("MONGODB_PASSWORD")

	credential := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoUser, mongoPass, mongoURL, mongoPort)

	return credential, nil
}
