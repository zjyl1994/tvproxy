package main

import (
	"net/http"
	"time"
)

func getHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}
