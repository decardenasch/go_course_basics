/*
Ejercicio 2 - Manipulando el body

Vamos a crear un endpoint llamado /saludo. Con una pequeña estructura con nombre y apellido que al pegarle
deberá responder en texto “Hola + nombre + apellido”

El endpoint deberá ser de método POST
Se deberá usar el package JSON para resolver el ejercicio
La respuesta deberá seguir esta estructura: “Hola Andrea Rivas”
La estructura deberá ser como esta:

	{
			“nombre”: “Andrea”,
			“apellido”: “Rivas”
	}
*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	type Person struct{
		Nombre string `json:"nombre"`
		Apellido string `json:"apellido"`
	}

	router.POST("/saludo",func(c *gin.Context){
		var person Person
		err := c.BindJSON(&person)
		if err != nil {
			c.JSON(http.StatusBadRequest,map[string]any{})
			return
		}
		c.JSON(http.StatusOK,map[string]any{
			"nombre" : "Andrea",
			"apellido" : "Rivas",
		})
		c.String(200,"Hola %s %s", person.Nombre, person.Apellido)

	})
	router.Run(":8080")
}