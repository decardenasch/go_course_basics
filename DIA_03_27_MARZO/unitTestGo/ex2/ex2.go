/*
	Ejercicio 2 - Calcular promedio
El colegio informó que las operaciones para calcular el promedio no se están realizando correctamente, por lo que ahora nos corresponde realizar los test correspondientes:
Calcular el promedio de las notas de los alumnos.
*/

package ex2 

import(
	"fmt"
)

func main(){
	fmt.Println(CalcularNotas(-1.0, 2.0, 3.0))
}

func CalcularNotas(notas ...float64) float64{
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
