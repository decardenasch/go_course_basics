/*
Ejercicio 2 - Imprimir datos

A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio. 
Ahora que el archivo sí existe, el panic no debe ser lanzado. 
Creamos el archivo “customers.txt” y le agregamos la información de los clientes. 
Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que 
contenga. En el caso de no poder leerlo, se debe lanzar un “panic”.
Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener 
un “defer” que nos indique que la ejecución finalizó. También recordemos cerrar los archivos al 
finalizar su uso.
*/

package main

import (
	"os"
	"fmt"
	"io"
)
func main(){
	fmt.Println("Iniciando...")
	defer func () {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Ejecución finalizada")
	}()
	// var f1 *os.File
	// f1, err := os.Create("customers.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// f1.WriteString("Nombre: Juan Perez")
	
	f1, err := os.Open("customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	file1, errRead := io.ReadAll(f1)

	if errRead != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

	fmt.Println(string(file1))
	defer f1.Close()
}