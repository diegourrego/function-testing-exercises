package exercise_1_test

import (
	"functions_exercises/exercise_1"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSalaryTaxes(t *testing.T) {
	t.Run("success - case 01: tax must be zero when salary is lower than $50.000", func(t *testing.T) {
		// Arrange
		testSalary := 27520.34
		expectedTax := 0.0
		// Act
		obtainedTax := exercise_1.SalaryTaxes(testSalary)
		// Assert
		require.Equal(t, expectedTax, obtainedTax)
	})

	t.Run("success - case 02: tax must be 17% when salary is up to $50.000 and less than $150.000",
		func(t *testing.T) {
			// Arrange
			testSalary := 57820.04
			expectedTax := testSalary * 0.17
			// Act
			obtainedTax := exercise_1.SalaryTaxes(testSalary)
			// Assert
			require.Equal(t, expectedTax, obtainedTax)
		})

	t.Run("success - case 02: tax must be 27% when salary is upper to $150.000", func(t *testing.T) {
		// Arrange
		testSalary := 163568.79
		expectedTax := testSalary * 0.27
		// Act
		obtainedTax := exercise_1.SalaryTaxes(testSalary)
		// Assert
		require.Equal(t, expectedTax, obtainedTax)
	})
}
