/* 
Ejercicio 4 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin. 

  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario: 
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.
*/

package main

import(
	"fmt"
	// "reflect" Sirve para saber el tipo de dato de una variable
)


func main(){
	company()
}

func company(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de Bejamin es",employees["Benjamin"]) //Acceder al elemento, teniendo la clave
	count := 0
	for key, element := range employees{
		if (key == "Benjamin"){
			fmt.Println("El nombre del empleado es",key ,"y su edad es", element)
		}
		if (element >21){
			count++
		}
	}
	fmt.Println("el numero de empleados mayores de 21 años es", count)
	employees["Federico"]= 25 //Agregar elementos al map
	fmt.Println(employees)
	delete(employees,"Pedro") //Eliminando a pedro del map
	fmt.Println(employees)
	
}

