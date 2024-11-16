package pets

type Cat struct {
	Animal
	Age    uint8
	Weight uint8
}

func (c *Cat) Eat(amount uint8) (uint8, error) {
	if amount > 5 {
		return 0, &ActionError{"Cat", "can't eat that much"}
	}
	return amount, nil
}

func (c *Cat) Walk() string {
	return "Cat walk"
}
