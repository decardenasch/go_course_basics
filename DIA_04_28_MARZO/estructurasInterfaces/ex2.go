/*
Ejercicio 2 - Productos
Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

Y los costos adicionales son:
Pequeño: solo tiene el costo del producto
Mediano: el precio del producto + un 3% de mantenerlo en la tienda
Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.
El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.
Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga el método Precio.

Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo del producto y los adicionales en caso que los tenga

*/

package main 

import (
	"fmt"
)

const (
	pequeño = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

func main() {
	producto,err :=factory(grande, 1000)
	if err != ""{
		fmt.Println(err)
		return
	}else{
		fmt.Println("Producto", grande)
		fmt.Println("Precio", producto.Precio())
	}
}
//Se crea la interfaz Producto
type Producto interface {
	//Se le añade el metodo Precio de tipo float64
	Precio() float64
}
//Se crean 3 estructuras que implementan la interfaz Producto
type productoPequeno struct {
	Costo float64
}
type productoMediano struct {
	Costo float64
}
type productoGrande struct {
	Costo float64
}
//Se crean los metodos Precio para cada estructura
func (p productoPequeno) Precio() float64{
	return p.Costo
}

func (p productoMediano) Precio() float64{
	return p.Costo + (p.Costo * 3 / 100)
}

func (p productoGrande) Precio() float64{
	return p.Costo + (p.Costo * 6 / 100) + 2500
}
//Se crea la funcion factory

func factory(tipo string, precio float64) (Producto,string){
	switch tipo {
	case pequeño:
		return productoPequeno{Costo: precio},""
	case mediano:
		return productoMediano{Costo: precio},""
	case grande:
		return productoGrande{Costo: precio},""
	default:
		return nil,"No existe el tipo de producto"
	}
}


