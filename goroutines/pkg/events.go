package event

import (
	"fmt"
	"time"
)

type Event struct {
	ID   int
	Name string
	Size int
}

func (e Event) String() string {
	return fmt.Sprintf("Event ID: %d, Name: %s", e.ID, e.Name)
}

func NewManager(interval time.Duration) *EventManager {
	return &EventManager{Interval: interval}
}

func NewProcessor() *EventProcessor {
	return &EventProcessor{}
}

type EventManager struct {
	Interval time.Duration
}

func (em *EventManager) Stream() <-chan Event {
	stream := make(chan Event)
	go func() {
		defer close(stream)
		for i := 1; i <= 1000; i++ {
			event := Event{
				ID:   i,
				Name: fmt.Sprintf("Event %d", i),
				Size: i * 100, // Adding a size field for demonstration
			}
			fmt.Printf("received event: %d | size: %d\n", event.ID, event.Size)
			stream <- event
			time.Sleep(em.Interval)
		}
	}()
	return stream
}

type EventProcessor struct{}

func (ep *EventProcessor) ProcessEvent(e Event) {
	fmt.Printf("Processing event: %s\n", e.String())
}
