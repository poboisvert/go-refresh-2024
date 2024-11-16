package operations

import (
	"sync"
)

var (
	//mu      sync.Mutex // go run . -race
	mu      sync.RWMutex // go run . -race
	balance int64
)

func Deposit(amount int64) {
	mu.Lock()
	//sema <- struct{}{}
	//<-sema
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int64 {
	mu.RLock()
	//sema <- struct{}{}
	//b := balance
	//<-sema
	defer mu.RUnlock()
	return balance
}

func Withdraw(amount int64) bool {
	mu.Lock()
	defer mu.Unlock()

	deposit(-amount)
	//sema <- struct{}{}
	//newBalance := balance - int64(amount)

	if balance < 0 {
		deposit(amount)
		return false
	}
	//balance = newBalance
	//<-sema
	return true
}

func deposit(amount int64) { balance += amount }
