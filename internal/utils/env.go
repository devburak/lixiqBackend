package utils

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MongoUri               string
	MongoDbName            string
	Port                   string
	AccessTokenPrivateKey  string
	AccessTokenPublicKey   string
	AccessTokenExpiredIn   time.Duration
	AccessTokenMaxAge      int
	RefreshTokenPrivateKey string
	RefreshTokenPublicKey  string
	RefreshTokenExpiredIn  time.Duration
	RefreshTokenMaxAge     int
}

func LoadConfig() *AppConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessTokenMaxAge := 15
	refreshTokenMaxAge := 60

	// Parse the duration values from environment variables
	accessTokenExpiredInStr := os.Getenv("ACCESS_TOKEN_EXPIRED_IN")
	refreshTokenExpiredInStr := os.Getenv("REFRESH_TOKEN_EXPIRED_IN")

	accessTokenExpiredIn, err := time.ParseDuration(accessTokenExpiredInStr)
	if err != nil {
		log.Fatal("Invalid value for ACCESS_TOKEN_EXPIRED_IN")
	}

	refreshTokenExpiredIn, err := time.ParseDuration(refreshTokenExpiredInStr)
	if err != nil {
		log.Fatal("Invalid value for REFRESH_TOKEN_EXPIRED_IN")
	}

	return &AppConfig{
		MongoUri:               os.Getenv("MONGO_URI"),
		MongoDbName:            os.Getenv("MONGO_DB_NAME"),
		Port:                   os.Getenv("PORT"),
		AccessTokenPrivateKey:  os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"),
		AccessTokenPublicKey:   os.Getenv("ACCESS_TOKEN_PUBLIC_KEY"),
		AccessTokenExpiredIn:   accessTokenExpiredIn,
		AccessTokenMaxAge:      accessTokenMaxAge,
		RefreshTokenPrivateKey: os.Getenv("REFRESH_TOKEN_PRIVATE_KEY"),
		RefreshTokenPublicKey:  os.Getenv("REFRESH_TOKEN_PUBLIC_KEY"),
		RefreshTokenExpiredIn:  refreshTokenExpiredIn,
		RefreshTokenMaxAge:     refreshTokenMaxAge,
	}
}
