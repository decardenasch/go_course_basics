/*
Ejercicio 2 - Clima
Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica 
de distintos lugares. 
Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión 
de donde te encuentres.
Imprimí los valores de las variables en consola.
¿Qué tipo de dato le asignarías a las variables?
*/
package main

import(
	"fmt"
)

func main(){
	var temperature float64 = 19.0
	var humidity int = 32
	var pressure int = 1009
// var (
// 	temperature float64 = 19.0
// 	humidity int = 32
// 	pressure int = 1009
// )
// var temperature,humidity,pressure float64 = 19.0,32.0,1009.0
	fmt.Println("La temperatura en medellin es" , temperature, "la humedad es" , humidity , "y la presion atmosferica es", pressure)
}

