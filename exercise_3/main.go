//EJERCICIOS FUNCIONES

package main

import (
	"fmt"
)

func main(){
	//fmt.Println(tax(54000))

}

// // Ejercicio 1 - Impuestos de salario
// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, 
//para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario. 
// Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más 
//de $150.000 se le descontará además un 10 % (27% en total).

func tax(salary float64)float64{
	if salary > 50000 && salary<=150000 {
		return salary*0.17
	}else if salary > 150000{
		return salary* 0.27
	}else {
		return salary*0
	}
}

// Ejercicio 2 - Calcular promedio

// Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.