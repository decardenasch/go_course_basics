/* 
Ejercicio 1 - Impuestos de salario #1
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje
 “Error: el salario ingresado no alcanza el mínimo imponible" y lanzalo en caso de que “salary” 
 sea menor a 150.000. De lo contrario, tendrás que imprimir por consola el mensaje
  “Debe pagar impuesto”.
*/
package main

import(
	"fmt"

)

func main() {
	var salary int
	salary = 100000
	err := validateSalary(salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Debe pagar impuesto")
}

type ErrorSalario struct {
	msg string
}

func (e *ErrorSalario) Error() string {
	return fmt.Sprintf("Error: %s",e.msg)
}

func validateSalary(salary int) error {
	if salary < 150000 {
		return &ErrorSalario{msg:"El salario ingresado no alcanza el mínimo imponible"}
	}
	return nil
}