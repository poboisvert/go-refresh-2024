package main

import "fmt"

func main() {

	var results []int
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Print("a\\b\\t")

	for _, a := range numbers {
		fmt.Printf("%8d", a)
	}
	fmt.Print("\n\n")

	for _, a := range numbers {
		fmt.Printf("%d\t\t", a)
		for _, b := range numbers {
			result := a * b
			results = append(results, result)
		}
	}

	for index, result := range results {
		fmt.Printf("result: %d\t%d\n", index, result)
	}
}
