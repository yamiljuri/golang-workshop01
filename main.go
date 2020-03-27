package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Estructura del tipo Persona
type Person struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

//Esta funcion se pasa como handler a la ruta que
//estamos creando con el endpoint "person"
//"localhost:8080/person", en este caso retornamos una
//structura de tipo Persona
func GetPerson(context *gin.Context){
	person := Person{"Lautaro","Diaz"}
	//Devolvemos en formato json a la peticion solicitada
	context.JSON(http.StatusOK,person)
}

//Esta funcion se pasa como handler a la ruta que
//estamos creando con el endpoint "meli"
//"localhost:8080/meli"
func GetMeli(context *gin.Context){
	//Devolvemos en formato json a la peticion solicitada
		context.JSON(http.StatusOK,gin.H{
			"message":"Bienvenido a Meli",
			"text":"Esta es otra forma de declarar la funcion",
		})
}


func main() {
	//Instanciamos el framework
	server := gin.Default()

	//Creamos la ruta /lagash por metodo GET
	//Ej localhost:8080/lagash
	server.GET("/lagash", func(context *gin.Context) {
		//Devolvemos en formato json a la peticion solicitada
		context.JSON(http.StatusOK,gin.H{
			"message":"Bienvenido a Go",
		})
	})

	//Creamos la ruta /lagash por metodo GET, con la diferencia que en esta ruta
	//le pasamos un parametro al que posteriormente lo obtenemos con el nombre de
	// "name", Ej localhost:8080/lagash/yamil
	server.GET("/lagash/:name", func(context *gin.Context) {
		//Extraemos el parametro "name" enviado por GET
		paramName := context.Param("name")
		//Devolvemos en formato json a la peticion solicitada ademas del mensaje
		//le agregamos al final el valor que pasamos en el parametro "name"
		context.JSON(http.StatusOK,gin.H{
			"message":"Bienvenido a Go "+paramName,
		})
	})

	//Creamos la ruta /lagash por metodo GET
	//Ej localhost:8080/meli
	server.GET("/meli", GetMeli)
	//Creamos la ruta /lagash por metodo GET
	//Ej localhost:8080/person
	server.GET("/person", GetPerson)

	// Inicio el servidor en el puerto 8080
	server.Run(":8080")
}