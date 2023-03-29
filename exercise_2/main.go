package main

import(
	"fmt"
	// "reflect" Sirve para saber el tipo de dato de una variable
)
// Ejercicio 1 - Letras de una palabra
// La Real Academia Española quiere saber cuántas letras tiene una palabra y 
// luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:
// Crear una aplicación que reciba  una variable con la palabra e imprimir la cantidad de letras 
// que contiene la misma.
// Luego, imprimí cada una de las letras.

func main(){
	// Spell("Amarillo")
	// Spell2("Amarillo")
	// loan(23,"Employee",150000,2) //Empleado que cumple todo
	// loan(15,"Employee",150000,2) //Empleado que no cumple edad
	// loan(24,"Unemployed",150000,2) //Empleado que no cumple empleo
	// loan(24,"Employee",10000,2) //Empleado que no cumple empleo
	// month(7)
	company()
}
// func Spell(word string){
// 	fmt.Println("La palabra tiene:", len(word), "letras")
// 	// fmt.Println(reflect.TypeOf(word))
// 	fmt.Println("Deletreando la palabra:")
// 	for i:=0; i < len(word) ; i++ {
// 		fmt.Println(string(word[i]))
// 	}
// }
// func Spell2(word string){
// 	count:= 0
// 	fmt.Println("Deletreando la palabra:")
// 	for index,items:= range word{
// 		count = index + 1
// 		fmt.Println(string(items)) //El tipo de dato que da aqui es una runa que es un tipo char en Go,
// 		//por lo que hay que convertirlo en string nuevamente.
// 	}
// 	fmt.Println("La palabra tiene: ", count , "letras")
// }

// Ejercicio 2 - Préstamo

// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello 
// tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes 
//cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. 
//Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000. 
// Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.
// Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
   					
// func loan(age int, employmentStatus string, salary float64, laborOld int){
// 	if (age < 22 || employmentStatus != "Employee" || laborOld < 1) {
// 		fmt.Println("Usted no cumple con las condiciones minimas para recibir un prestamo")
// 	}else{
// 		fmt.Println("Usted si cumple con las condiciones minimas para recibir un prestamo")
// 			if salary > 100000 {
// 				fmt.Println("Adicionalmente el banco no le cobrará interes")
// 			}else{
// 				fmt.Println("Sin embargo el banco si le cobrará interes")
// 			}
// 	}
// }

// Ejercicio 3 - A qué mes corresponde

// Realizar una aplicación que reciba  una variable con el número del mes. 
// Según el número, imprimir el mes que corresponda en texto. 
// ¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
// Ej: 7, Julio.
// Nota: Validar que el número del mes, sea correcto.

// func month(num int){
// 	if (num < 0 || num > 12){
// 		fmt.Println("El valor del numero del mes no es correcto")
// 	}else{
// 		monthMap := map[int]string{1:"Enero",2:"Febrero",3:"Marzo",4:"Abril",5:"Mayo",6:"Junio",7:"Julio",8:"Agosto",9:"Septiembre",10:"Octubre",11:"Novimebte",12:"Diciembre"}
// 		for key, element := range monthMap{
// 			if (key == num){
// 				fmt.Println(num,",",element)
// 			}
// 		}	
// 	}
// }

// Ejercicio 4 - Qué edad tiene...
// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin. 

//   var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

// Por otro lado también es necesario: 
// Saber cuántos de sus empleados son mayores de 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del mapa.

func company(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de Bejamin es",employees["Benjamin"]) //Acceder al elemento, teniendo la clave
	count := 0
	for key, element := range employees{
		if (key == "Benjamin"){
			fmt.Println("El nombre del empleado es",key ,"y su edad es", element)
		}
		if (element >21){
			count++
		}
	}
	fmt.Println("el numero de empleados mayores de 21 años es", count)
	employees["Federico"]= 25 //Agregar elementos al map
	fmt.Println(employees)
	delete(employees,"Pedro") //Eliminando a pedro del map
	fmt.Println(employees)
	
}

