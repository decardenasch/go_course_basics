/*
Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. 
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y 
devuelva el promedio. No se pueden introducir notas negativas.

*/

package main 

import(
	"fmt"
)

func main(){
	fmt.Println(calcularNotas(-1.0, 2.0, 3.0))
}
func calcularNotas(notas ...float64) float64{
	suma := 0.0
	for _,items:= range notas {
		if items < 0 {
			fmt.Println("No acepta notas negativas")
			break
		}else {
			suma += items
		}
	}
	promedio := suma/ float64(len(notas))
	return promedio
}