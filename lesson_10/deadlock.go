package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Millisecond)
		//ch <- val
		ch <- 66
		fmt.Println("Received: 66")
	}()

	select {
	case r := <-ch:
		fmt.Println(r)
	case <-time.After(2 * time.Second):
		fmt.Println("Empty or invalid data")
	}

}
