package pets

type Dog struct {
	Animal
	Age    uint8
	Weight uint8
}

func (c *Dog) Eat(amount uint8) (uint8, error) {
	if amount > 20 {
		return 0, &ActionError{"Dog", "can't eat that much"}
	}

	return amount, nil
}

func (c *Dog) Walk() string {
	return "Dog walk"
}
