package ex4

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T){
	//Arrange : Preparas el escenario
	t.Run("Average", func(t *testing.T) {
		//Arrange : Preparas el escenario
		args := []int{10, 9, 8, 8, 4}
		expectedResult := 4.0
		//Act : Pones hacer al programa lo que tiene que hacer
		obtainedResult:= AverageFunc(args...)
		//Assert : Verificacion
		assert.Equal(t,expectedResult,obtainedResult)
	})
	t.Run("Minimum", func(t *testing.T) {
		//Arrange : Preparas el escenario
		args := []int{10, 9, 8, 8, 4}
		expectedResult := 7.8
		//Act : Pones hacer al programa lo que tiene que hacer
		obtainedResult:= AverageFunc(args...)
		//Assert : Verificacion
		assert.Equal(t,expectedResult,obtainedResult)
	})
	t.Run("Maximum", func(t *testing.T) {
		//Arrange : Preparas el escenario
		args := []int{10, 9, 8, 8, 4}
		expectedResult := 10.0
		//Act : Pones hacer al programa lo que tiene que hacer
		obtainedResult:= MaxFunc(args...)
		//Assert : Verificacion
		assert.Equal(t,expectedResult,obtainedResult)
	})
}
