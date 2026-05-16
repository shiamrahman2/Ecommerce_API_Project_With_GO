package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration Config

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
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
	jwtSecretKey:=os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey==""{
		fmt.Println("Give Valid Secret Key")
		os.Exit(1)
	}
	configuration = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpport,
		JwtSecretKey:jwtSecretKey,
	}
}

func GetConfig() Config {
	loadConfig()
	return configuration
}
