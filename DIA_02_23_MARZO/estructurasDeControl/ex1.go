/*
Ejercicio 1 - Letras de una palabra
La Real Academia Española quiere saber cuántas letras tiene una palabra y 
luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:
Crear una aplicación que reciba  una variable con la palabra e imprimir la cantidad de letras 
que contiene la misma.
Luego, imprimí cada una de las letras.
*/
package main

import(
	"fmt"
)

func main(){
	Spell("Amarillo")
	Spell2("Amarillo")

 }

 func Spell(word string){
	fmt.Println("La palabra tiene:", len(word), "letras")
	// fmt.Println(reflect.TypeOf(word))
	fmt.Println("Deletreando la palabra:")
	for i:=0; i < len(word) ; i++ {
		fmt.Println(string(word[i]))
	}
}
func Spell2(word string){
	count:= 0
	fmt.Println("Deletreando la palabra:")
	for index,items:= range word{
		count = index + 1
		fmt.Println(string(items)) //El tipo de dato que da aqui es una runa que es un tipo char en Go,
		//por lo que hay que convertirlo en string nuevamente.
	}
	fmt.Println("La palabra tiene: ", count , "letras")
}

