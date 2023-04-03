/*
Ejercicio 1 - Prueba de Ping
Vamos a crear una aplicación Web con el framework Gin que tenga un endpoint /ping que al pegarle responda un texto que diga “pong”
El endpoint deberá ser de método GET
La respuesta de “pong” deberá ser enviada como texto, NO como JSON
*/
package main 

import (
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context){
		// El primer parametro del metodo String es el codigo de estado HTTP, y el segundo 
		//es la cadena que se desee
		c.String(200,"pong")
	})
	router.Run(":8080")	
}