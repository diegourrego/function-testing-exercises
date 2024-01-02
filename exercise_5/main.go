package exercise_5

func Animal(animal string) (func(animalsQuantity uint8) float32, string) {
	switch animal {
	case "dog":
		return calculateDogsFood, ""
	case "cat":
		return calculateCatsFood, ""
	case "hamster":
		return calculateHamstersFood, ""
	case "tarantula":
		return calculateTarantulasFood, ""
	default:
		return nil, "animal not found. check it and try again."
	}
}

func calculateDogsFood(animalsQuantity uint8) float32 {
	return float32(animalsQuantity) * 10.0
}

func calculateCatsFood(animalsQuantity uint8) float32 {
	return float32(animalsQuantity) * 5.0
}

func calculateHamstersFood(animalsQuantity uint8) float32 {
	return float32(animalsQuantity) * 0.25
}

func calculateTarantulasFood(animalsQuantity uint8) float32 {
	return float32(animalsQuantity) * 0.15
}
