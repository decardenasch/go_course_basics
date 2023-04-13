/*
Ejercicio 1
Crear un programa que cumpla los siguiente puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.   
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado. 
Ejecutar al menos una vez cada método y función definido desde main().

*/

package main 

import (
	"fmt"
)
// Se crea la estructura llamada Product con los campos ID, Name, Price, Description y Category.
type Product struct{
	ID 			int
	Name 		string
	Price 		float64
	Description string
	Category	string
}
//Se crea un slice global de Product llamado Products instanciado con valores.
var Products = []Product{{1,"Pepsi",1.5,"Bebida","Bebidas"}}

func main() {
	/* Para llamar un metodo primero se debe crear una instancia de la estructura y luego llamar
	al metodo en esa instancia
	*/
	//Se instancia un producto para guardar en el slice
	p2 := Product{2,"Coca Cola",1.5,"Bebida","Bebidas"}
	p3 := Product{3,"Manzana Postobon",4.5,"Bebida","Bebidas"}
	//Se utiliza el método Save() para guardar el producto en el slice
	p2.Save()
	p3.Save()
	fmt.Println("--------------Show Products-----------")
	//Se utiliza el método GetAll() para mostrar todos los productos del slice
	p3.GetAll()
	fmt.Println("--------------Search Products-----------")
	//Se utiliza la función getById() para buscar un producto por ID
	producto, err := getById(3)
	//Se valida si el producto existe o no
	if err == "" {
		fmt.Println(producto)
	} else {
		fmt.Println(err)
	}	
}

func (p Product) Save(){
	//Añade el producto desde el cual se llama al método
	Products = append(Products, p)
}

func (p Product) GetAll(){
	//Recorre el slice de productos y los muestra
	for _, items := range Products{
		fmt.Println(items)
	}
}
//Funcion que busca por ID
func getById(id int) (Product,string){
	for _, items := range Products{
		//Comparamos el id ingresado con los ID de todo el slice
		if items.ID == id{
			return items,""
		}
	}
	return Product{},"No se encontro el producto"
}