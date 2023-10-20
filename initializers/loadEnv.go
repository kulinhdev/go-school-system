package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	DBUser                 string
	DBPassword             string
	DBName                 string
	DBHost                 string
	DBPort                 string
	RunPort                string
	AccessTokenExpiresIn   time.Duration
	AccessTokenPrivateKey  string
	AccessTokenPublicKey   string
	RefreshTokenExpiresIn  time.Duration
	RefreshTokenPrivateKey string
	RefreshTokenPublicKey  string
	AccessTokenMaxAge      int
	RefreshTokenMaxAge     int
}

var Config *AppConfig

func Initialize() {
	// Load the .env file.
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	// Read environment variables or use default values.
	Config = &AppConfig{
		DBUser:                 os.Getenv("DB_USER"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBName:                 os.Getenv("DB_NAME"),
		DBHost:                 os.Getenv("DB_HOST"),
		DBPort:                 os.Getenv("DB_PORT"),
		RunPort:                os.Getenv("RUN_PORT"),
		AccessTokenExpiresIn:   time.Second * time.Duration(parseEnvInt("ACCESS_TOKEN_EXPIRES_IN", 3600)),
		AccessTokenPrivateKey:  os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"),
		AccessTokenPublicKey:   os.Getenv("ACCESS_TOKEN_PUBLIC_KEY"),
		AccessTokenMaxAge:      parseEnvInt("ACCESS_TOKEN_MAX_AGE", 86400),
		RefreshTokenExpiresIn:  time.Second * time.Duration(parseEnvInt("REFRESH_TOKEN_EXPIRES_IN", 86400)),
		RefreshTokenPrivateKey: os.Getenv("REFRESH_TOKEN_PRIVATE_KEY"),
		RefreshTokenPublicKey:  os.Getenv("REFRESH_TOKEN_PUBLIC_KEY"),
		RefreshTokenMaxAge:     parseEnvInt("REFRESH_TOKEN_MAX_AGE", 86400),
	}
}

func parseEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Printf("Invalid value for %s, using the default: %v\n", key, defaultValue)
		return defaultValue
	}
	return value
}

func GetDatabaseURL() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", Config.DBHost, Config.DBUser, Config.DBPassword, Config.DBName, Config.DBPort)
}
