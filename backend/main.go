package main

import (
	"fmt"
	"os"

	"github.com/tocoteron/kankaku/infrastructure/web"
)

// Get environment var or default value
func getEnvOrDefault(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

// Get environment var or panic program
func getEnvOrPanic(key string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	panic(fmt.Sprintf("%s is not set", key))
}

func main() {
	// Get env vars
	port := getEnvOrDefault("PORT", "8080")
	secret := []byte(getEnvOrPanic("SECRET"))

	// Start server
	server := web.NewServer(port, secret)
	server.Run()
}
