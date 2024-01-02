package main

import (
	"fmt"
	"functions_exercises/exercise_1"
	"functions_exercises/exercise_2"
	"functions_exercises/exercise_3"
)

func main() {
	// Exercise #1:
	exercise_1.SalaryTaxes(150001)

	// Exercise #2:
	average, err := exercise_2.CalculateAverage(4.3, 2.1, 3.0)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Printf("The average of notes is: %v\n", average)

	// Exercise #3:
	employeeSalary := exercise_3.CalculateSalary(160, "C")
	fmt.Printf("Tehe employee salary is: %v\n", employeeSalary)
}
