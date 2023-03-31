/*
Ejercicio 2 - Préstamo

Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello 
tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes 
cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. 
Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000. 
Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.
Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/
package main

import(
	"fmt"
)
func main(){
	loan(23,"Employee",150000,2) //Empleado que cumple todo
	// loan(15,"Employee",150000,2) //Empleado que no cumple edad
	// loan(24,"Unemployed",150000,2) //Empleado que no cumple empleo
	// loan(24,"Employee",10000,2) //Empleado que no cumple empleo

}		

func loan(age int, employmentStatus string, salary float64, laborOld int){
	if (age < 22 || employmentStatus != "Employee" || laborOld < 1) {
		fmt.Println("Usted no cumple con las condiciones minimas para recibir un prestamo")
	}else{
		fmt.Println("Usted si cumple con las condiciones minimas para recibir un prestamo")
			if salary > 100000 {
				fmt.Println("Adicionalmente el banco no le cobrará interes")
			}else{
				fmt.Println("Sin embargo el banco si le cobrará interes")
			}
	}
}
