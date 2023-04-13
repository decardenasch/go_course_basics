/*
Ejercicio 5 - Calcular cantidad de alimento
El refugio de animales envió una queja ya que el cálculo total de alimento a comprar no fue el correcto y compraron menos alimento del que necesitaban. Para mantener satisfecho a nuestro cliente deberemos realizar los test necesarios para verificar que todo funcione correctamente:
Verificar el cálculo de la cantidad de alimento para los perros.
Verificar el cálculo de la cantidad de alimento para los gatos.
Verificar el cálculo de la cantidad de alimento para los hamster.
Verificar el cálculo de la cantidad de alimento para las tarántulas.

func TestDog(t *testing.T)
func TestCat(t *testing.T)

*/
package ex5 

// import (
// 	"fmt"
// )
const (
	dog    = "dog"
	cat    = "cat"
	hamster    = "hamster"
	tarantula    = "tarantula"
 )

// func main(){
// 	fncAnimal, msg := animal("tarantula")
// 	if msg != ""{
// 		fmt.Println(msg)
// 	}
// 	fmt.Println(fncAnimal(5))

// }
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