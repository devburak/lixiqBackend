package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	ConnectionPoolSize     int
	EmailHost              string
	EmailFrom              string
	EmailPassword          string
	EmailPort              int
}

func LoadConfig() *AppConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessTokenMaxAge, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_MAXAGE"))
	if err != nil {
		log.Fatal("Invalid value for ACCESS_TOKEN_MAXAGE")
	}
	refreshTokenMaxAge, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_MAXAGE"))
	if err != nil {
		log.Fatal("Invalid value for REFRESH_TOKEN_MAXAGE")
	}
	connectionPoolSize, err := strconv.Atoi(os.Getenv("CONNECTION_POOL_SIZE"))
	if err != nil {
		log.Fatal("Invalid value for connection pool size")
	}

	// Parse the duration values from environment variables
	accessTokenExpiredInStr := os.Getenv("ACCESS_TOKEN_EXPIRED_IN")
	refreshTokenExpiredInStr := os.Getenv("REFRESH_TOKEN_EXPIRED_IN")

	accessTokenExpiredIn, err := time.ParseDuration(accessTokenExpiredInStr)
	fmt.Println("access token expired in", accessTokenExpiredIn)
	if err != nil {
		log.Fatal("Invalid value for ACCESS_TOKEN_EXPIRED_IN")
	}

	refreshTokenExpiredIn, err := time.ParseDuration(refreshTokenExpiredInStr)
	fmt.Println("refresh token expired in", refreshTokenExpiredIn)
	if err != nil {
		log.Fatal("Invalid value for REFRESH_TOKEN_EXPIRED_IN")
	}

	emailHost := os.Getenv("E_MAIL_HOST")
	emailFrom := os.Getenv("E_MAIL_FROM")
	emailPassword := os.Getenv("E_MAIL_PASSWORD")
	emailPort, err := strconv.Atoi(os.Getenv("E_MAIL_PORT"))
	if err != nil {
		log.Fatal("Invalid value for E_MAIL_PORT")
	}

	return &AppConfig{
		MongoUri:               os.Getenv("MONGODB_LOCAL_URI"),
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
		ConnectionPoolSize:     connectionPoolSize,
		EmailHost:              emailHost,
		EmailFrom:              emailFrom,
		EmailPassword:          emailPassword,
		EmailPort:              emailPort,
	}
}
