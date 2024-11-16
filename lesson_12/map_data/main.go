package main

import (
	"fmt"
	loader "map_data/data"
	"sync"
)

// var allData = make(map[string]int)
var allData = map[string]int{
	"a":   loader.LoadData("a"),
	"www": loader.LoadData("ww"),
	"e":   loader.LoadData("e"),
	"y":   loader.LoadData("y"),
}

func main() {

	var wg sync.WaitGroup

	names := []string{"q", "u", "e", "t", "y"}

	for _, name := range names {
		wg.Add(1)

		go func(name string) {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				fmt.Printf("Goroutine for %s: Icon value = %d\n", name, Data(name))
			}
		}(name)
	}
	wg.Wait()
}

/* func Data(name string) int {
	if data, ok := allData[name]; ok {
		return data
	}

	data := loader.LoadData(name)
	allData[name] = data
	return data
} */

func Data(name string) int {
	return allData[name]
}
