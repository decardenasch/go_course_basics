/* 
Ejercicio 2 - Impuestos de salario #2
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: 
el salario es menor a 10.000" y lanzalo en caso de que “salary” sea menor o igual a  10.000.
 La validación debe ser hecha con la función “Is()” dentro del “main”.
*/
package main

import(
	"fmt"
	"errors"
)

func main() {
	var salary int
	salary = 9000
	err := validateSalary(salary)
	fmt.Println(errors.Is(err, ErrSalary))
	if err != nil {
		if errors.Is(err, ErrSalary){
			fmt.Println(err)
			return

		}
	}
	fmt.Println("Debe pagar impuesto")
}
type ErrorSalario struct {
	msg string
}

func (e *ErrorSalario) Error() string {
	return fmt.Sprintf("Error: %s",e.msg)
}

var ErrSalary = &ErrorSalario{msg:"El salario ingresado no alcanza el mínimo imponible"}

func validateSalary(salary int) error {
	if salary < 10000 {
		return ErrSalary
	}
	return nil
}