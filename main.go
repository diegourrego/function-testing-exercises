package main

import (
	"fmt"
	"functions_exercises/exercise_1"
	"functions_exercises/exercise_2"
)

func main() {
	exercise_1.SalaryTaxes(150001)
	average, err := exercise_2.CalculateAverage(4.3, 2.1, 3.0)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Printf("The average of notes is: %v", average)
}
