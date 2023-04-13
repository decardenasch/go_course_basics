/* 
Ejercicio 5 -  Impuestos de salario #5
Vamos a hacer que nuestro programa sea un poco más complejo y útil. 
Desarrollá las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar 
el 10 % en concepto de impuesto. En caso de que la cantidad de horas mensuales ingresadas sea 
menor a 80 o un número negativo, la función debe devolver un error. El mismo tendrá que indicar 
“Error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
*/
package main

import(
	"fmt"
	"errors"
)

func main(){
	var horas int
	var valorHora float64
	horas = 80
	valorHora = 19000
	salarioCalculado, err := SalarioMensual(horas, valorHora)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("El salario calculado es de: %.2f", salarioCalculado)
}

func SalarioMensual (horas int, valorHora float64)(salarioCalculado float64, err error){
	if horas < 80 {
		err = errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return
	}
	salarioCalculado = float64(horas) * valorHora
	if salarioCalculado >= 150000 {
		salarioCalculado = salarioCalculado - (salarioCalculado * 0.1)
	}
	return
}