package main

import (
	"os"

	"github.com/nmatsui/bearer-auth-api/router"
)

const DEFAULT_PORT = "8080"

func main() {
	port := os.Getenv("LISTEN_PORT")
	if len(port) == 0 {
		port = DEFAULT_PORT
	}

	handler := router.NewHandler()
	handler.Run(":" + port)
}
