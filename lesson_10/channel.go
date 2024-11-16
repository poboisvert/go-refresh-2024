package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// Unbuffered Channel
// func main() {
// 	channel := make(chan int)

// 	go func() {
// 		channel <- 666
// 		fmt.Println("Sent:666")
// 	}()

// 	val := <-channel
// 	fmt.Println("Received:", val)
// }

func request(server string) string {
	delay := rand.Intn(1000)

	fmt.Printf("Requesting %s (will take %d ms) \n", server, delay)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("Response received from %s\n", server)

	return fmt.Sprintf("Response from %s", server)
}

func mirroredQuery() string {
	response := make(chan string, 2)

	go func() { response <- request("www.espn.com") }()
	go func() { response <- request("www.sportingnews.com") }()
	go func() { response <- request("www.cbssports.com") }()
	go func() { response <- request("www.foxsports.com") }()
	go func() { response <- request("www.nfl.com") }()
	go func() { response <- request("www.nba.com") }()
	go func() { response <- request("www.mlb.com") }()
	go func() { response <- request("www.nhl.com") }()
	go func() { response <- request("www.mls.com") }()
	go func() { response <- request("www.pga.com") }()

	return <-response
}

// Buffered Channel. Add limiter
func main() {
	// channel := make(chan int, 2)

	// go func() {
	//	channel <- 666
	//	channel <- 266
	//	channel <- 166
	//	fmt.Println("Sent:666")
	// }()

	result := mirroredQuery()
	fmt.Println("Final result in main(): ", result)

	time.Sleep(2 * time.Second)

	fmt.Printf("Number of goroutines live: %d\n", runtime.NumGoroutine())
	//val := <-channel
	//fmt.Println("Received:", val)
}
