/* 
Ejercicio 3 - Impuestos de salario #3

Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, 
en reemplazo de “Error()”,  se implemente “errors.New()”.
*/
package main
import(
	"fmt"
	"errors"
)

func main() {
	var salary int
	salary = 19000
	err := validateSalary(salary)

	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("Debe pagar impuesto")
	}
}


func validateSalary(salary int) error {
	if salary < 10000 {
		return errors.New("El salario ingresado no alcanza el mínimo imponible")
	}
	return nil
}