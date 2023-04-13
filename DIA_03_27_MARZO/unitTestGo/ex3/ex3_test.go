package ex3

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculoSalario(t *testing.T)  {
	t.Run("Salario de la categoria A", func(t *testing.T) {
		//Arrange : Preparas el escenario
		minutos := 120.0
		categoria := "A"
		expectedResult := 9000.0
		//Act : Pones hacer al programa lo que tiene que hacer
		obtainedResult := CalculoSalario(minutos, categoria)
		//Assert : Verificacion
		assert.Equal(t, expectedResult, obtainedResult, "el salario obtenido no es el esperado")
	})
	t.Run("Salario de la categoria B", func(t *testing.T) {
		//Arrange : Preparas el escenario
		minutos := 120.0
		categoria := "B"
		expectedResult := 4200.0
		//Act : Pones hacer al programa lo que tiene que hacer
		obtainedResult := CalculoSalario(minutos, categoria)
		//Assert : Verificacion
		assert.Equal(t, expectedResult, obtainedResult, "el salario obtenido no es el esperado")
	})
	t.Run("Salario de la categoria C", func(t *testing.T) {
		//Arrange : Preparas el escenario
		minutos := 120.0
		categoria := "C"
		expectedResult := 2000.0
		//Act : Pones hacer al programa lo que tiene que hacer
		obtainedResult := CalculoSalario(minutos, categoria)
		//Assert : Verificacion
		assert.Equal(t, expectedResult, obtainedResult, "el salario obtenido no es el esperado")
	})
	
}