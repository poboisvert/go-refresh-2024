// Package main demonstrates concurrent event processing with rate limiting
package main

import (
	"time"

	events "goroutines/pkg"
)

func main() {
	// Create an event manager that generates events every 1/10th of a second
	manager := events.NewManager(time.Second / 300)
	// Create an event processor to handle the events
	processor := events.NewProcessor()
	// Create a channel witNewProcessorh buffer size 8 to limit concurrent goroutines
	limiter := make(chan int, 10)

	// Continuously receive events from the manager's stream
	for event := range manager.Stream() {
		// Block if we already have 8 goroutines running
		limiter <- 10
		// Launch a new goroutine to process each event
		go func(e events.Event) {
			processor.ProcessEvent(e)
			// Release the limiter slot when done
			<-limiter
		}(event)
	}
}
