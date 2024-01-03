package exercise_4

import "math"

func CalculateStatistics(operationType string) (func(nums ...int) int, string) {

	switch operationType {
	case "minimum":
		return findMinimum, ""
	case "average":
		return calculateAverage, ""
	case "maximum":
		return findMaximum, ""
	default:
		return nil, "Unexpected operation type. Please check it again and retry"
	}
}

func findMaximum(nums ...int) int {
	var maxNumber int
	for _, num := range nums {
		if num > maxNumber {
			maxNumber = num
		}
	}
	return maxNumber
}

func findMinimum(nums ...int) int {
	var minNumber int = math.MaxInt
	for _, num := range nums {
		if num < minNumber {
			minNumber = num
		}
	}
	return minNumber
}

func calculateAverage(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum / len(nums)
}
