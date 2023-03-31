/*
Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.

*/

package main 

import (
	"fmt"
)

func main (){
	fmt.Println("El salario es:", calculoSalario(120,"A"))

}

func calculoSalario (minutos float64, categoria string) float64{
	var horas float64 = minutos / 60;
	switch categoria {
		case "A":
			return  horas * 3000 + (horas * 3000 * 0.5)
		case "B":
			return horas * 1500 + (horas * 3000 * 0.2)
		case "C":
			return horas * 1000
	}
	return 0
}