package main

import (
	"net/http"
	"time"
	"log"
	"os"
	"io"
)

func main() {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// HTTP request
	response, err := client.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	
	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)
}

