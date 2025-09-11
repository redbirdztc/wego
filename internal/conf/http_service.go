package conf

import (
	"fmt"
	"os"
)

func GetPort() string {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}
	if !validatePort(port) {
		panic(fmt.Errorf("invalid port: %s", port))
	}
	return port
}

func validatePort(port string) bool {
	// Check if port is between valid range 1-65535
	if port == "" {
		return false
	}

	portNum := 0
	_, err := fmt.Sscanf(port, "%d", &portNum)
	if err != nil {
		return false
	}

	return portNum > 0 && portNum <= 65535
}
