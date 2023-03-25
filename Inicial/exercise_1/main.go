package main

import(
	"fmt"
)
// Ejercicio 1 - Imprimí tu nombre
	// Tendrás que: 
	// Crear una aplicación donde tengas como variable tu nombre y dirección.
	// Imprimir en consola el valor de cada una de las variables.

func main(){
 	// name := "Lorena"
 	// dir := "Carrera 56A # 61 - 24"
 	var name string = "Lorena"
 	var dir string = "Carrera 56A # 61 - 24"
	fmt.Println("Mi nombre es: " , name)
 	fmt.Println("Mi direccion es: ", dir)
 }

//  Ejercicio 2 - Clima
	// Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica 
	//de distintos lugares. 
	// Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión 
	//de donde te encuentres.
	// Imprimí los valores de las variables en consola.
	// ¿Qué tipo de dato le asignarías a las variables?

// func main(){
// 	var temperature float64 = 19.0
// 	var humidity int = 32
// 	var pressure int = 1009
// // var (
// // 	temperature float64 = 19.0
// // 	humidity int = 32
// // 	pressure int = 1009
// // )
// // var temperature,humidity,pressure float64 = 19.0,32.0,1009.0
// 	fmt.Println("La temperatura en medellin es" , temperature, "la humedad es" , humidity , "y la presion atmosferica es", pressure)

// }

//  Ejercicio 3 - Declaración de variables
	// Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia 
	// Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos 
	// del examen consiste en declarar distintas variables. 
	// Necesita ayuda para:
	// Detectar cuáles de estas variables que declaró el alumno son correctas.
	// Corregir las incorrectas.

// var 1nombre string //Incorrecta no se puede nombrar una variable con un numero al inicio -> var nombre string
//   var apellido string //Correcta
//   var int edad //Incorrecta primero viene el nombre y luego el tipo -> var edad int
//   1apellido := 6 //Incorrecta no se puede nombrar una variable con un numero al inicio, 
//   nombre de la variable debe ser algo que corresponda a lo que tiene guardado, en este caso seria un string 
//   -> apellido := "6"
//   var licencia_de_conducir = true //Incorrecto las variables deben llamarse en camelCase -> var licenciaDeConducir = true
//   var estatura de la persona int //Incorrecto las variables no deben tener espacio en su nombre -> var 
//   estaturaDeLaPersona int
//   cantidadDeHijos := 2 //Correcto con formato lowercamelCase

//   Ejercicio 4 - Tipos de datos
	//   Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos 
	//   tipos en Go, pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó su código 
	//   y pidió la ayuda de un desarrollador experimentado que pueda:
	//   Verificar su código y realizar las correcciones necesarias.
  
// 	var apellido string = "Gomez" //Correcto
// 	var edad int = "35" //El tipo de dato no corresponde -> var edad int = 35
// 	boolean := "false"; //Correcto
// 	var sueldo string = 45857.90 //El tipo de dato no corresponde -> var sueldo string = "45857.90" 
// 	var nombre string = "Julián" //Correcto
  