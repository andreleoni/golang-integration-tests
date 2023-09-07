package main

import (
	"log"
	"net/http"

	"gopkg.in/resty.v1"
)

type KeepAliveInterface interface {
	KeepAlive(url string) bool
}

func main() {
	client := resty.New()

	resp, _ := client.R().Get("https://www.google.com.br")

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("Error on GET: %v", resp.StatusCode())
	} else {
		log.Printf("Success response of request!")
	}
}
