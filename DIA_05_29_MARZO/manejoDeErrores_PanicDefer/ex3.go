/*
Ejercicio 3 - Registrando clientes
El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos 
de nuevos clientes. Los datos requeridos son:
Legajo
Nombre 
DNI
Número de teléfono
Domicilio


Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe. Para ello, necesitás 
leer los datos de un array. En caso de que esté repetido, debes manipular adecuadamente el error 
como hemos visto hasta aquí. Ese error deberá:
1.- generar un panic;
2.- lanzar por consola el mensaje: “Error: el cliente ya existe”, y continuar con la ejecución 
del programa normalmente.
Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función 
para validar que todos los datos a registrar de un cliente contienen un valor distinto de cero. 
Esta función debe retornar, al menos, dos valores. Uno de ellos tendrá que ser del tipo error 
para el caso de que se ingrese por parámetro algún valor cero (recordá los valores cero de cada 
	tipo de dato, ej: 0, “”, nil).
Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por 
consola los siguientes mensajes: “Fin de la ejecución” y “Se detectaron varios errores en tiempo 
de ejecución”. Utilizá defer para cumplir con este requerimiento.

Requerimientos generales:
Utilizá recover para recuperar el valor de los panics que puedan surgir
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error. 
Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de 
Go (realiza también la validación pertinente para el caso de error retornado).

*/

package main

import (
	"fmt"

)

var Clientes = []Cliente{}

type Cliente struct {
	Legajo    int
	Nombre    string
	DNI	      int
	Telefono  int
	Domicilio string
}

func main(){
	c1 := Cliente{1, "Juan", 123456789, 123456789, "Calle falsa 123"}
	c2 := Cliente{2, "Pedro", 123456789, 123456789, "Calle falsa 123"}
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("Fin de la ejecución")	
	}()
	guardarCliente(c1)
	guardarCliente(c2)
	fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	fmt.Println(Clientes)
	
}

//Comprobar que existe un cliente
func cExist(dni int)(bool,error){
	for _, cliente := range Clientes{
		if cliente.DNI == dni{
			return true, fmt.Errorf("Error: el cliente ya existe")
		}
	}
	return false, nil
}
func cValidate(c Cliente)(bool,error){
	if (c.Legajo == 0 || c.Nombre == "" || c.DNI == 0 || c.Telefono == 0 || c.Domicilio == ""){
		return false, fmt.Errorf("Error: el cliente no tiene todos los datos")
	}
	return true, nil
}

func guardarCliente(c Cliente){
	if _,err := cExist(c.DNI); err != nil{
		panic(err)
	}
	if _, err := cValidate(c); err != nil{
		panic(err)
	}
	Clientes = append(Clientes, c)
}

	