package main

import (
	"fmt"
	"functions_exercises/exercise_1"
	"functions_exercises/exercise_2"
	"functions_exercises/exercise_3"
	"functions_exercises/exercise_4"
)

func main() {
	// Exercise #1:
	exercise_1.SalaryTaxes(150001)

	// Exercise #2:
	averageOfNotes, err := exercise_2.CalculateAverage(4.3, 2.1, 3.0)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Printf("The average of notes is: %v\n", averageOfNotes)

	// Exercise #3:
	employeeSalary := exercise_3.CalculateSalary(160, "C")
	fmt.Printf("Tehe employee salary is: %v\n", employeeSalary)

	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	)

	minFunc, err := exercise_4.CalculateStatistics(minimum)
	if err != "" {
		fmt.Println("ERROR. Please check the argument")
	}

	maxFunc, err := exercise_4.CalculateStatistics(maximum)
	if err != "" {
		fmt.Println("ERROR. Please check the argument")
	}

	averageFunc, err := exercise_4.CalculateStatistics(average)
	if err != "" {
		fmt.Println("ERROR. Please check the argument")
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)         // 2
	averageValue := averageFunc(2, 3, 3, 4, 10, 2, 4, 5) // 10
	maxValue := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)         // 4

	fmt.Printf("The min value is: %v, The max value is: %v and the average is: %v\n",
		minValue, maxValue, averageValue)

}
