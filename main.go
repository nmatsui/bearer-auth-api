package main

import (
	"os"
	"strconv"

	"github.com/nmatsui/bearer-auth-api/router"
)

const LISTEN_PORT = "LISTEN_PORT"
const DEFAULT_PORT = "8080"

func main() {
	handler := router.NewHandler()
	handler.Run(getListenPort())
}

func getListenPort() string {
	port := os.Getenv(LISTEN_PORT)
	if len(port) == 0 {
		port = DEFAULT_PORT
	}
	intPort, err := strconv.Atoi(port)
	if err != nil || intPort < 1 || 65535 < intPort {
		port = DEFAULT_PORT
	}

	return ":" + port
}
