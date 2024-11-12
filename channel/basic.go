package main

import "fmt"

func main() {
	dataChan := make(chan int, 2) // must add space into it else impossible

	// goroutine add memory

	dataChan <- 789
	dataChan <- 123

	n := <-dataChan // GET data from channel
	fmt.Printf("n = %d\n", n)

	n = <-dataChan // GET data from channel second value from int, 2
	fmt.Printf("n = %d\n", n)
}
