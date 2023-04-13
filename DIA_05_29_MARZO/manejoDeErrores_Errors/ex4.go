/* 
Ejercicio 4 - Impuestos de salario #4
Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error 
reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible 
(el mensaje mostrado por consola deberá decir: “Error: el mínimo imponible es de 150.000 y 
el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro). 
*/
package main

import(
	"fmt"

)

func main() {
	var salary int
	salary = 190000
	err := validateSalary(salary)

	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("Debe pagar impuesto")
	}
}


func validateSalary(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("Error: el minimo imponible es de 150.000 y el salario ingresado es de : %d",salary) 
	}
	return nil
}
