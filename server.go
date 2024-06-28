package main

import (
	"net/http"
	"os"
)

func StartServer() {
	http.HandleFunc("/callback", WebhookHandler())

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "80"
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Fatal(err)
	}
}
