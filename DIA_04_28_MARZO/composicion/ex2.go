/*
Ejercicio 2 - Employee
Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará a gestionar correctamente dichos empleados. Los objetivos son:
Crear una estructura Person con los campos ID, Name, DateOfBirth.
Crear una estructura Employee con los campos: ID, Position y una composicion con la estructura Person.
Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión de los campos de un empleado.
Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar el método PrintEmployee().
Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.


*/

package main 

import("fmt")

type Person struct{
	ID 			int
	Name 		string
	DateOfBirth string
}
//Embedding una estructura dentro de otra estructura
type Employee struct{
	ID 			int
	Position 	string
	Person 		Person
}
//Asignar un metodo a una estructura que tiene embedding
func (e Employee) PrintEmployee(){
	fmt.Printf("%d,%s,%d,%s,%s\n",e.ID,e.Position,e.Person.ID,e.Person.Name,e.Person.DateOfBirth)
}
func main(){
	//Instanciando la primera estructura
	p1 := Person{
		ID: 1,
		Name: "Juan",
		DateOfBirth: "12/12/1990",
	}
	//Instanciando la estructura embedding
	e1 := Employee{
		ID: 1,
		Position: "Programador",
		Person: p1}
	//Instanciando la estructura embedding de otra forma
	// e1 := Employee{
	// 	ID: 1,
	// 	Position: "Programador",
	// 	Person: Person{
	// 		ID: 1,
	// 		Name: "Juan",
	// 		DateOfBirth: "12/12/1990",
	// 	},
	// }
	//Llamando el metodo de la estructura embedding
	e1.PrintEmployee()
}