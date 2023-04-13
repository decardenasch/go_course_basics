/*
Ejercicio 4 - Testear el cálculo de estadísticas
Los profesores de la universidad de Colombia, entraron al programa de análisis de datos  de Google, el cual premia a los mejores estadísticos de la región. Por ello los profesores nos solicitaron comprobar el correcto funcionamiento de nuestros cálculos estadísticos. Se solicita la siguiente tarea:
Realizar test para calcular el mínimo de calificaciones.
Realizar test para calcular el máximo de calificaciones.
Realizar test para calcular el promedio de calificaciones.


*/

package ex4 

import(
	//"fmt"
	"math"
	"errors"
)
const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
	)

// func main(){
// 	const (
// 		minimum = "minimum"
// 		average = "average"
// 		maximum = "maximum"
// 	 )
// 	 operationFinal, err := Operation(average)
// 	 if err != nil {
// 		fmt.Println(err)
// 	 }else{
// 		fmt.Println(operationFinal(2, 3, 3, 4, 10, 2, 4, 5))
// 	 }
// 	// minFunc, err := operation(minimum)
// 	// averageFunc, err := operation(average)
// 	// maxFunc, err := operation(maximum)
// 	// minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
// 	// fmt.Printf("El valor minimo es: %d \n", minValue )

// 	// averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
// 	// fmt.Println("El promedio es: ", averageValue )

// 	// maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
// 	// fmt.Printf("El valor maximo es: %d \n", maxValue )

// }

func MinFunc (marks ...int) float64{
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

func AverageFunc (marks ...int) float64{
	sum := 0.0
	for _,items := range marks{
		sum += float64(items)
	}
	return sum/float64(len(marks))
}

func MaxFunc (marks ...int) float64{
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

func Operation(typeOperation string) (func(marks ...int) float64, error){
	switch typeOperation{
		case "minimum":
			return MinFunc,nil
		case "average":
			return AverageFunc,nil
		case "maximum":
			return MaxFunc,nil
		default:
			return nil, errors.New("Ha ingresado una opcion no valida")
	} 

}
