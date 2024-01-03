package exercise_4_test

import (
	"functions_exercises/exercise_4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateStatistics(t *testing.T) {
	t.Run("success - case01: Must show the minimum number correctly and return an empty error message",
		func(t *testing.T) {
			// Arrange
			var mockNumbers = []int{1, 7, 8, 23, 54, 2}
			var mockOperationType = "minimum"
			var minimumNumberExpected = 1

			minimumNumberFunction, err := exercise_4.CalculateStatistics(mockOperationType)

			// Act
			minimumNumberObtained := minimumNumberFunction(mockNumbers...)
			// Assert
			require.Equal(t, minimumNumberExpected, minimumNumberObtained)
			require.Equal(t, "", err)

		})

	t.Run("success - case02: Must show the maximum number correctly and return an empty error message",
		func(t *testing.T) {
			// Arrange
			var mockNumbers = []int{1, 7, 8, 23, 54, 2}
			var mockOperationType = "maximum"
			var minimumNumberExpected = 54

			maximumNumberFunction, err := exercise_4.CalculateStatistics(mockOperationType)

			// Act
			maximumNumberObtained := maximumNumberFunction(mockNumbers...)
			// Assert
			require.Equal(t, minimumNumberExpected, maximumNumberObtained)
			require.Equal(t, "", err)

		})
	t.Run("success - case03: Must show the average correctly and return an empty error message",
		func(t *testing.T) {
			// Arrange
			var mockNumbers = []int{1, 7, 8, 23, 54, 2}
			var mockOperationType = "average"
			var averageExpected = (1 + 7 + 8 + 23 + 54 + 2) / 6

			averageFunction, err := exercise_4.CalculateStatistics(mockOperationType)

			// Act
			averageObtained := averageFunction(mockNumbers...)
			// Assert
			require.Equal(t, averageExpected, averageObtained)
			require.Equal(t, "", err)

		})
	t.Run("Failure - case01: Must return an nil and an error message when operation type is incorrect",
		func(t *testing.T) {
			// Arrange
			var mockOperationType = "other"

			returnedFunction, err := exercise_4.CalculateStatistics(mockOperationType)

			// Assert
			require.Nil(t, returnedFunction, "Expected nil function")
			require.Equal(t, "Unexpected operation type. Please check it again and retry", err)

		})
}
