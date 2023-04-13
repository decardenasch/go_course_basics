package ex1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// func TestTax(t *testing.T) {
// 	//Arrange : Preparas el escenario
// 	salary := 45000.0
// 	expectedResult := 0.0
// 	//Act : Pones hacer al programa lo que tiene que hacer
// 	obtainedResult := Tax(salary)
// 	//Assert : Verificacion
// 	if expectedResult != obtainedResult {
// 		t.Errorf("Expected %f, got %f", expectedResult, obtainedResult)
// 	}
// }

func TestTax(t *testing.T){
	t.Run("Salario por debajo de $50000",func(t *testing.T){
		//Arrange : Preparas el escenario
		salary:= 45000.0
		expectedResult := 0.0
		//Act : Pones hacer al programa lo que tiene que hacer
		obtainedResult := Tax(salary)
		//Assert : Verificacion
		assert.Equal(t,expectedResult, obtainedResult,"el impuesto obtenido no es el esperado")
	})
	t.Run("Salario entre $50000 y $150000",func(t *testing.T){
		//Arrange : Preparas el escenario
		salary:= 54000.0
		expectedResult := 9180.0
		//Act : Pones hacer al programa lo que tiene que hacer
		obtainedResult := Tax(salary)
		//Assert : Verificacion
		assert.Equal(t,expectedResult, obtainedResult,"el impuesto obtenido no es el esperado")
	})
	t.Run("Salario por encima de $150000",func(t *testing.T){
		//Arrange : Preparas el escenario
		salary:= 200000.0
		expectedResult := 54000.0
		//Act : Pones hacer al programa lo que tiene que hacer
		obtainedResult := Tax(salary)
		//Assert : Verificacion
		assert.Equal(t,expectedResult, obtainedResult,"el impuesto obtenido no es el esperado")
	})
}