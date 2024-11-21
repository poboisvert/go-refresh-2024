package test

import (
	"sync/atomic"
	"testing"
)

/*
"Atomic operation" is a term in computer science used to describe an operation or a group of operations
that cannot be divided. This means that during the execution of an atomic operation,
the operation either fully executes or does not execute at all,
and there are no intermediate states or partial executions.

Atomic operations are commonly used in multi-threaded programming and concurrency control to ensure data consistency and reliability. In a multi-threaded environment,
multiple threads may simultaneously access and modify shared data.
Without proper synchronization and control of these operations,
data races and unpredictable behavior can occur.
Atomic operations provide an effective way to handle such concurrent access and ensure data correctness.

The benefits of using atomic operations include:
1. **Data Integrity**: Ensures that shared data is updated correctly without corruption.
2. **Performance**: Atomic operations can be faster than using locks, as they avoid the overhead of acquiring and releasing locks.
3. **Simplicity**: They simplify the code by reducing the need for complex locking mechanisms.

To demonstrate the difference between a non-atomic operation and an atomic operation, we have two scenarios below.

Scenario 1 (normal test):
1) Access go_practice/test
2) Run: go test -run TestDataRaceCondition
*/
func TestDataRaceCondition(t *testing.T) {
	var counter int32
	for i := 0; i < 10; i++ {
		go func(i int) {
			counter += int32(i) // This can lead to a race condition
		}(i)
	}
}

/*
Atomic scenario 2 (race condition test):
1) Access go_practice/test
2) Run: go test -run TestAtomicDataRaceCondition -race
*/
func TestAtomicDataRaceCondition(t *testing.T) {
	var counter int32
	for i := 0; i < 10; i++ {
		go func(i int) {
			atomic.AddInt32(&counter, int32(i)) // Safely increment the counter atomically
		}(i)
	}
}

// An alternative asynchronous solution using channels to avoid race conditions
func TestAsyncSolution(t *testing.T) {
	var counter int32
	ch := make(chan int32)

	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- int32(i) // Send the value to the channel
		}(i)
	}

	for i := 0; i < 10; i++ {
		counter += <-ch // Receive the value from the channel and update the counter
	}
}
