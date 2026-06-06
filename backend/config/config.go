package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	Env       string
	DBUrl     string
	JWTSecret string
}

var config *Config

func GetConfig() *Config {
	return config
}


func Load() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	env := os.Getenv("ENV")
	if env == "" || (env != "production" && env != "testing" && env != "remote_testing") {
		env = "development"
	}

	config = &Config{
		Port:      getEnv("PORT", "PORT_BACKUP"),
		Env:       env,
		DBUrl:     buildDbUrl(env),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	if config.Port == "" {
		config.Port = getEnv("PORT_BACKUP", "3030")
	}
}

func buildDbUrl(env string) string {
	if env == "production" {
		return os.Getenv("DB_URL")
	}

	if env == "remote_testing" {
		return os.Getenv("DB_URL_testing")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	
	dbName := os.Getenv("DB_DEV_NAME")
	if env == "testing" {
		dbName = os.Getenv("DB_TEST_NAME")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}