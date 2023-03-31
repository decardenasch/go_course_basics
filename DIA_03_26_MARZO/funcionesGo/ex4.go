/*
Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso. 
Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva 
otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y devuelva el cálculo 
que se indicó en la función anterior.
Ejemplo:

...
*/

package main 

import(
	"fmt"
	"math"
	"errors"
)

func main(){
	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	 )
	 operationFinal, err := operation(average)
	 if err != nil {
		fmt.Println(err)
	 }else{
		fmt.Println(operationFinal(2, 3, 3, 4, 10, 2, 4, 5))
	 }
	// minFunc, err := operation(minimum)
	// averageFunc, err := operation(average)
	// maxFunc, err := operation(maximum)
	// minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	// fmt.Printf("El valor minimo es: %d \n", minValue )

	// averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	// fmt.Println("El promedio es: ", averageValue )

	// maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	// fmt.Printf("El valor maximo es: %d \n", maxValue )

}

func minFunc (marks ...int) float64{
	//Se selecciona el valor maximo entero que existe
	//Se recorre marks y se va asignando el valor menor que se va encontrando
	min := math.MaxInt
	for _,items := range marks{
		if items < min{
			min = items
		}
	}
	return float64(min)
}
func averageFunc (marks ...int) float64{
	sum := 0.0
	for _,items := range marks{
		sum += float64(items)
	}
	return sum/float64(len(marks))
}

func maxFunc (marks ...int) float64{
	//Se selecciona math como el numero mas pequeño negativo que existe para un entero
	//Se recorre marks y se va asignando el valor mayor que se va encontrando
	max := int(math.Inf(-1))
	for _, items:= range marks {
        if items > max {
            max = items
        }
    }
    
    return float64(max)
}

func operation(typeOperation string) (func(marks ...int) float64, error){
	switch typeOperation{
		case "minimum":
			return minFunc,nil
		case "average":
			return averageFunc,nil
		case "maximum":
			return maxFunc,nil
		default:
			return nil, errors.New("Ha ingresado una opcion no valida")
	} 

}
