# go-refresh-2024

Go project 2024

In Go, concurrency is achieved through the use of goroutines, channels, and mutexes.

**Goroutines**: A goroutine is a lightweight thread that can run concurrently with other goroutines. For example, in a web server, each incoming request can be handled in a separate goroutine, allowing the server to handle multiple requests simultaneously.

Example:

```go
package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(i)
	}
}

func main() {
	go printNumbers()
	printLetters()
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}
```

In this example, `printNumbers` and `printLetters` are two goroutines that run concurrently, printing numbers and letters to the console.

**Channels**: A channel is a communication mechanism that allows goroutines to send and receive data. For example, in a producer-consumer scenario, a producer goroutine can send data to a channel, and a consumer goroutine can receive and process the data.

Example:

```go
package main

import (
	"fmt"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
```

In this example, the `producer` goroutine sends numbers to the channel, and the `consumer` goroutine receives and prints the numbers.

**Mutexes**: A mutex (short for "mutual exclusion") is a mechanism that allows only one goroutine to access a shared resource at a time. For example, in a bank account, a mutex can be used to ensure that only one goroutine can deposit or withdraw money at a time, preventing race conditions.

Example:

```go
package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (a *BankAccount) deposit(amount int) {
	a.mu.Lock()
	a.balance += amount
	a.mu.Unlock()
}

func (a *BankAccount) withdraw(amount int) {
	a.mu.Lock()
	if a.balance >= amount {
		a.balance -= amount
	}
	a.mu.Unlock()
}

func main() {
	account := &BankAccount{balance: 100}
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		account.deposit(50)
		wg.Done()
	}()
	go func() {
		account.withdraw(20)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(account.balance)
}
```

In this example, the `BankAccount` struct uses a mutex to ensure that only one goroutine can access the `balance` field at a time.

**-race param**: The `-race` flag is a command-line option for the Go compiler that enables the race detector. The race detector is a tool that helps identify race conditions in Go programs. For example, if two goroutines are trying to access the same variable at the same time, the race detector will detect this and report it as a potential race condition.

To use the `-race` flag, simply add it to your `go run` or `go build` command:

```bash
go run -race main.go
```

This will enable the race detector and report any potential race conditions in your program.
