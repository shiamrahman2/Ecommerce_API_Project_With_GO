package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Name          string
	Host          string
	Password      string
	Port          int
	User          string
	EnableSSLMode bool
}

var configuration *Config

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("ENV file can't Read ", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")

	if serviceName == "" {
		fmt.Println("Service Name Required")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HTTP PORT is required")
		os.Exit(1)
	}
	httpport, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("HTTP PORT must be Number")
		os.Exit(1)
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Give Valid Secret Key")
		os.Exit(1)
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB Host Required")
		os.Exit(1)
	}
	dbPort := os.Getenv("DB_PORT")
	dbPrt, err := strconv.Atoi(dbPort)
	if err != nil {
		fmt.Println("Port must be interger")
		os.Exit(1)
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB Password is Required")
		os.Exit(1)
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB User is Required")
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB Name is Required")
		os.Exit(1)
	}
	dbSSLMode := os.Getenv("DB_ENABLE_SSL_MODE")
	dbBoolSSLMode, err := strconv.ParseBool(dbSSLMode)
	if err != nil {
		fmt.Println("DB SSL MODE must be boolean")
		os.Exit(1)
	}
	dbConfig := &DBConfig{
		Name:          dbName,
		Host:          dbHost,
		Password:      dbPassword,
		User:          dbUser,
		Port:          dbPrt,
		EnableSSLMode: dbBoolSSLMode,
	}
	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     httpport,
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}

}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
