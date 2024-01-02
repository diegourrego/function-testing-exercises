package main

import (
	"fmt"
	"functions_exercises/exercise_1"
	"functions_exercises/exercise_2"
	"functions_exercises/exercise_3"
	"functions_exercises/exercise_4"
	"functions_exercises/exercise_5"
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

	const (
		dog       = "dog"
		cat       = "cat"
		hamster   = "hamster"
		tarantula = "tarantula"
	)

	dogsFoodFunc, msg := exercise_5.Animal(dog)

	if msg != "" {
		fmt.Println(msg)
	}

	catsFoodFunc, msg := exercise_5.Animal(cat)

	if msg != "" {
		fmt.Println(msg)
	}

	hamstersFoodFunc, msg := exercise_5.Animal(hamster)

	if msg != "" {
		fmt.Println(msg)
	}

	tarantulasFoodFunc, msg := exercise_5.Animal(tarantula)

	if msg != "" {
		fmt.Println(msg)
	}

	var amount float32
	amount += dogsFoodFunc(10)
	amount += catsFoodFunc(10)
	amount += hamstersFoodFunc(10)
	amount += tarantulasFoodFunc(10)

	fmt.Printf("The final amount for the food is: %v kg\n", amount)

}
