package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetDataSourceUrl() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := getEnvironmentValue("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Couldn't convert supplied port %v to a valid integer", portStr)
	}
	return port
}

func getEnvironmentValue(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Required env var %s missing, exiting!", key)
	}
	return value
}
