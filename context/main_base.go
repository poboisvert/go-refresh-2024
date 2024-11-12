package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	// Add timeout context
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "https://placehold.co/600x400", nil)
	if err != nil {
		panic(err)
	}

	// Perform HTTP GET
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	imageData, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("download img size: %d\n", len(imageData))
}
