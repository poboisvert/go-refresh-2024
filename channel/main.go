package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulating expensive calculation
func Sleep() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func main() {
	dataChan := make(chan int) // must add space into it else impossible

	go func() {
		// Create a WaitGroup to wait for all goroutines to finish
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1) // Increment the WaitGroup counter
			go func() {
				defer wg.Done()    // Decrement the counter when the goroutine completes
				result := Sleep()  // Simulate an expensive calculation
				dataChan <- result // Send the result to the channel
			}()
		}
		wg.Wait()       // Wait for all goroutines to finish
		close(dataChan) // Close the channel to signal that no more values will be sent
	}()
	// goroutine add memory
	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}

}
