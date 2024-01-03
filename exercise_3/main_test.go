package exercise_3_test

import (
	"functions_exercises/exercise_3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateSalary(t *testing.T) {
	t.Run("Success - case01: Must calculate the salary for 'A' category successfully",
		func(t *testing.T) {
			// Arrange
			var mockCategory string = "A"
			var mockBaseSalary float32 = 3000
			var mockHoursPerMonth float32 = 120
			var mockSalaryBonus = (mockBaseSalary * 8 * 30) * 0.5

			expectedSalary := (mockBaseSalary * mockHoursPerMonth) + mockSalaryBonus

			// Act
			obtainedSalary, err := exercise_3.CalculateSalary(mockHoursPerMonth, mockCategory)
			//Assert
			require.Equal(t, expectedSalary, obtainedSalary)
			require.Equal(t, "", err)
		})

	t.Run("Success - case02: Must calculate the salary for 'B' category successfully",
		func(t *testing.T) {
			// Arrange
			var mockCategory string = "B"
			var mockBaseSalary float32 = 1500
			var mockHoursPerMonth float32 = 150
			var mockSalaryBonus = (mockBaseSalary * 8 * 30) * 0.2

			expectedSalary := (mockBaseSalary * mockHoursPerMonth) + mockSalaryBonus

			// Act
			obtainedSalary, err := exercise_3.CalculateSalary(mockHoursPerMonth, mockCategory)
			//Assert
			require.Equal(t, expectedSalary, obtainedSalary)
			require.Equal(t, "", err)
		})

	t.Run("Success - case03: Must calculate the salary for 'C' category successfully",
		func(t *testing.T) {
			// Arrange
			var mockCategory string = "C"
			var mockBaseSalary float32 = 1000
			var mockHoursPerMonth float32 = 160

			expectedSalary := mockBaseSalary * mockHoursPerMonth

			// Act
			obtainedSalary, err := exercise_3.CalculateSalary(mockHoursPerMonth, mockCategory)
			//Assert
			require.Equal(t, expectedSalary, obtainedSalary)
			require.Equal(t, "", err)
		})

	t.Run("Failure - case01: Must return an error for an invalid category",
		func(t *testing.T) {
			// Arrange
			var mockCategory string = "D"
			var mockHoursPerMonth float32 = 10

			// Act
			obtainedSalary, err := exercise_3.CalculateSalary(mockHoursPerMonth, mockCategory)
			//Assert
			require.Equal(t, float32(0.0), obtainedSalary)
			require.Equal(t, "Unexpected category. Check it and try again.", err)
		})
}
