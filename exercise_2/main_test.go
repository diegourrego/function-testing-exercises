package exercise_2_test

import (
	"functions_exercises/exercise_2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	t.Run("success - case01: Must do the notes average successfully", func(t *testing.T) {
		// Arrange
		mockNotes := []float64{2.3, 4.7, 5.0, 1.2, 3.01}
		expectedAverage := (2.3 + 4.7 + 5.0 + 1.2 + 3.01) / 5
		// Act
		obtainedAverage, err := exercise_2.CalculateAverage(mockNotes...)
		// Assert
		require.Equal(t, expectedAverage, obtainedAverage)
		require.Equal(t, err, "")
	})

	t.Run("failure - case01: Must show an error when gets a negative note and a zero average", func(t *testing.T) {
		// Arrange
		mockNotes := []float64{2.3, -1.2, 5.0, 1.2, 3.01}
		// Act
		obtainedAverage, err := exercise_2.CalculateAverage(mockNotes...)
		// Assert
		require.Equal(t, 0.0, obtainedAverage)
		require.Equal(t, err, "ERROR!. You can't introduce negative notes")
	})
}
