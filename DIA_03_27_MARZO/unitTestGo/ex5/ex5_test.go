package ex5

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAnimal (t *testing.T){
	t.Run("TestDog", func(t *testing.T){
		//Arrange
		amountAnimal := 5
		expectedResult := 50.0
		//Act
		obtainedResult:= amountFoodDog(amountAnimal)
		//Assert
		assert.Equal(t,expectedResult,obtainedResult)
	})
	t.Run("TestCat", func(t *testing.T){
		//Arrange
		amountAnimal := 5
		expectedResult := 25.0
		//Act
		obtainedResult:= amountFoodCat(amountAnimal)
		//Assert
		assert.Equal(t,expectedResult,obtainedResult)
	})
	t.Run("TestHamster", func(t *testing.T){
		//Arrange
		amountAnimal := 5
		expectedResult := 1.25
		//Act
		obtainedResult:= amountFoodHamster(amountAnimal)
		//Assert
		assert.Equal(t,expectedResult,obtainedResult)
	})
	t.Run("TestTarantula", func(t *testing.T){
		//Arrange
		amountAnimal := 5
		expectedResult := 0.75
		//Act
		obtainedResult:= amountFoodTarantula(amountAnimal)
		//Assert
		assert.Equal(t,expectedResult,obtainedResult)
	})
}