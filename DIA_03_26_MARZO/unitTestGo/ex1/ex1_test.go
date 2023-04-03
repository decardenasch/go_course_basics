package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTax(t *testing.T){
	//Arrange : Preparas el escenario
	salary:= 45000.0
	expectedResult := 0.0
	//Act : Pones hacer al programa lo que tiene que hacer
	obtainedResult := Tax(salary)
	//Assert : Verificacion 
	if expectedResult != obtainedResult {
		t.Errorf("Expected %f, got %f", expectedResult, obtainedResult)
	}
}

func TestTax(t *testing.T){
	//Arrange : Preparas el escenario
	salary:= 45000.0
	expectedResult := 0.0
	//Act : Pones hacer al programa lo que tiene que hacer
	obtainedResult := Tax(salary)
	//Assert : Verificacion 
	assert.Equal(t,expectedResult, obtainedResult)
}