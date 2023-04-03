/*
Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, 
hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.
Tienen los siguientes datos:
Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.
Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función 
y un mensaje (en caso que no exista el animal).
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
Ejemplo:


...


animalDog, msg := animal(dog)
animalCat, msg := animal(cat)


...


var amount float64
amount += animalDog(10)
amount += animalCat(10)


*/

package main 

import (
	"fmt"
)
const (
	dog    = "dog"
	cat    = "cat"
	hamster    = "hamster"
	tarantula    = "tarantula"
 )

func main(){
	fncAnimal, msg := animal("tarantula")
	if msg != ""{
		fmt.Println(msg)
	}
	fmt.Println(fncAnimal(5))

}
func amountFoodDog(amountAnimal int) float64{
	return float64(amountAnimal) * 10
}
func amountFoodCat(amountAnimal int) float64{
	return float64(amountAnimal) * 5
}
func amountFoodHamster(amountAnimal int) float64{
	return float64(amountAnimal) * 0.250
}
func amountFoodTarantula(amountAnimal int) float64{
	return float64(amountAnimal) * 0.150
}
func animal (typeAnimal string) (func(amountFood int) float64, string){
	switch typeAnimal{
		case dog:
			return amountFoodDog,"" //Pasa la funcion de referencia, no la llama por lo que no se le pasa los (), queda ahi lista para luego usar.
		case cat:
			return amountFoodCat, ""
		case hamster:
			return amountFoodHamster, ""
		case tarantula:
			return amountFoodTarantula,""
		default: 
			return nil, "No tenemos este animal. Esperamos poder darle refugio pronto."
	}
}