/*
	La  empresa de chocolates que anteriormente necesitaba calcular el impuesto de sus empleados, 
	al momento de depositar el sueldo de los mismos ahora nos solicitó validar que los cálculos de estos 
	impuestos están correctos. Para esto nos encargaron el trabajo de realizar los test correspondientes para:
	Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
	Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
	Calcular el impuesto en caso de que el empleado gane por encima de $150.000.

*/

package main 


import (
	"fmt"
)

func main(){
	fmt.Println(tax(54000))

}

func tax(salary float64)float64{
	if salary > 50000 && salary<=150000 {
		return salary*0.17
	}else if salary > 150000{
		return salary* 0.27
	}else {
		return salary*0
	}
}
