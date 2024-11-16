package main

import (
	"animals/pets"
	"fmt"
	"os"
)

func main() {
	oneCat := pets.Cat{
		Animal: pets.Animal{Name: "Mr. Lulu"},
		Age:    10}

	oneDog := pets.Dog{
		Animal: pets.Animal{Name: "Mr. Lulu"},
		Age:    1,
		Weight: 4}

	var unitFeedToCat uint8 = 3
	var unitFeedToDog uint8 = 10

	catAction, err := feed(&oneCat, unitFeedToCat)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error feeding cat: %v\n", err)
	} else {
		fmt.Println("Cat ate: ", catAction)
	}

	fmt.Print("\n\n\t ==== \n\n\n")

	dogAction, err := feed(&oneDog, unitFeedToDog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error feeding cat: %v\n", err)
	} else {
		fmt.Println("Cat ate: ", dogAction)
	}
}

func feed(animal pets.EatWalk, amount uint8) (uint8, error) {
	switch v := animal.(type) {
	case *pets.Cat:
		fmt.Println(v.GetName(), "is a cat aged", v.Age)
	case *pets.Dog:
		fmt.Println(v.GetName(), "is a dog aged", v.Age)
	default:
		fmt.Println("Unknown animal type")
	}

	fmt.Println("First, let's walk!")
	fmt.Println(animal.Walk())

	fmt.Println("Now, let's feed our", animal.GetName())

	return animal.Eat(amount)
}

func displayInfo(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("This is a string:", v)
	case int:
		fmt.Println("This is an int:", v)
	case pets.Cat:
		fmt.Println("This is a Cat named:", v.Name, "and it is", v.Age, "years old")
	case pets.Dog:
		fmt.Println("This is a Dog named:", v.Name, "and weight:", v.Weight)
	default:
		fmt.Println("Unknown type")
	}
}
