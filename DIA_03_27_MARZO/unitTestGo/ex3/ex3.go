/*
Ejercicio 3 - Test del salario
La empresa marinera no está de acuerdo con los resultados obtenidos en los cálculos de los salarios, por ello nos piden realizar una serie de tests sobre nuestro programa. Necesitaremos realizar las siguientes pruebas en nuestro código:
Calcular el salario de la categoría “A”.
Calcular el salario de la categoría “B”.
Calcular el salario de la categoría “C”.

*/

package ex3 

// import (
// 	"fmt"
// )

// func main (){
// 	fmt.Println("El salario es:", CalculoSalario(120,"A"))

// }

func CalculoSalario (minutos float64, categoria string) float64{
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