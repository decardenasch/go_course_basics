package ex2

import (
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestCalcularNotas(t *testing.T) {
	t.Run("Promedio de notas", func(t *testing.T) {
		//Arrange : Preparas el escenario
		notas := []float64{3.0, 4.0, 5.0}
		expectedResult := 4.0
		//Act : Pones hacer al programa lo que tiene que hacer
		obtainedResult := CalcularNotas(notas...)
		//Assert : Verificacion
		assert.Equal(t, expectedResult, obtainedResult, "el promedio obtenido no es el esperado")
	})
}
