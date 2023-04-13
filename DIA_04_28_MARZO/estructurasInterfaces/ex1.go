/*
Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle


*/

package main 
import (
	"fmt")

type Alumnos struct{
	Nombre 		string
	Apellido 	string
	DNI 		int
	Fecha 		string
}

var Alumno = []Alumnos{}
func main(){
	// Crear un alumno con los datos de un alumno/a
	a := Alumnos{"Juan", "Perez", 12345678, "12/12/2020"}
	a1 := Alumnos{"Maria", "Gomez", 87654321, "12/12/2020"}
	a2 := Alumnos{"Pedro", "Gomez", 87654321, "12/12/2020"}
	a.Save()
	a1.Save()
	a2.Save()
	a.Detalle()

}

func (a Alumnos) Save() {
	// Guardar los datos del alumno en un archivo
	Alumno = append(Alumno, a)
}
func (a Alumnos) Detalle(){
	// Mostrar todos los alumnos registrados
	for _, items := range Alumno{
		fmt.Printf("Nombre: %s\n Apellido: %s\n DNI: %d\n Fecha: %s\n", items.Nombre, items.Apellido, items.DNI, items.Fecha)
	}
}