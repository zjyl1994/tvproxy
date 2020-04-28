package main

import (
	"net/http"
	"time"
)

func getHttpClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}
