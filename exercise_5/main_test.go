package exercise_5_test

import (
	"functions_exercises/exercise_5"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDog(t *testing.T) {
	// Arrange
	mockAnimal := "dog"
	var mockDogsQuantity uint8 = 4
	expectedDogFood := float32(mockDogsQuantity) * 10.0

	// Act
	calculateDogFoodFunc, err := exercise_5.Animal(mockAnimal)
	dogFoodObtained := calculateDogFoodFunc(mockDogsQuantity)
	// Assert
	require.Equal(t, expectedDogFood, dogFoodObtained)
	require.Equal(t, "", err)
}

func TestCat(t *testing.T) {
	// Arrange
	mockAnimal := "cat"
	var mockDogsQuantity uint8 = 7
	expectedDogFood := float32(mockDogsQuantity) * 5.0

	// Act
	calculateDogFoodFunc, err := exercise_5.Animal(mockAnimal)
	dogFoodObtained := calculateDogFoodFunc(mockDogsQuantity)
	// Assert
	require.Equal(t, expectedDogFood, dogFoodObtained)
	require.Equal(t, "", err)
}

func TestHamster(t *testing.T) {
	// Arrange
	mockAnimal := "hamster"
	var mockDogsQuantity uint8 = 12
	expectedDogFood := float32(mockDogsQuantity) * 0.25

	// Act
	calculateDogFoodFunc, err := exercise_5.Animal(mockAnimal)
	dogFoodObtained := calculateDogFoodFunc(mockDogsQuantity)
	// Assert
	require.Equal(t, expectedDogFood, dogFoodObtained)
	require.Equal(t, "", err)
}

func TestTarantula(t *testing.T) {
	// Arrange
	mockAnimal := "tarantula"
	var mockDogsQuantity uint8 = 2
	expectedDogFood := float32(mockDogsQuantity) * 0.15

	// Act
	calculateDogFoodFunc, err := exercise_5.Animal(mockAnimal)
	dogFoodObtained := calculateDogFoodFunc(mockDogsQuantity)
	// Assert
	require.Equal(t, expectedDogFood, dogFoodObtained)
	require.Equal(t, "", err)
}

func TestAnimal_Error(t *testing.T) {
	// Arrange
	mockAnimal := "horse"

	// Act
	returnedFunction, err := exercise_5.Animal(mockAnimal)
	// Assert
	require.Nil(t, returnedFunction, "Should return a nil instead of a function")
	require.Equal(t, "animal not found. check it and try again.", err)
}
