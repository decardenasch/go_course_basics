/*
Ejercicio 3 - A qué mes corresponde

Realizar una aplicación que reciba  una variable con el número del mes. 
Según el número, imprimir el mes que corresponda en texto. 
¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.
Nota: Validar que el número del mes, sea correcto.
*/

package main

import(
	"fmt"
	// "reflect" Sirve para saber el tipo de dato de una variable
)

func main(){
	month(7)
}
func month(num int){
	if (num < 0 || num > 12){
		fmt.Println("El valor del numero del mes no es correcto")
	}else{
		monthMap := map[int]string{1:"Enero",2:"Febrero",3:"Marzo",4:"Abril",5:"Mayo",6:"Junio",7:"Julio",8:"Agosto",9:"Septiembre",10:"Octubre",11:"Novimebte",12:"Diciembre"}
		for key, element := range monthMap{
			if (key == num){
				fmt.Println(num,",",element)
			}
		}	
	}
}

