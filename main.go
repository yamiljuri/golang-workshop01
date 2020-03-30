package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

const (
	CREDENTIAL_USERNAME = "lagash"
	CREDENTIAL_PASSWORD = "meli"
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

//En esta funcion hicimos un metodo de autentication customizado
// devolvemos una funcion que recibe como parametro el contexto
func basicAuth(username string , password string) gin.HandlerFunc{
	return func(context *gin.Context) {
		auth := strings.SplitN(context.Request.Header.Get("Authorization")," ",2)
		//Comprobamos si viene el Header "Authorization" y si es de tipo basic
		// basic se forma por Basic user:name (user:name tiene que ir en base64)
		if len(auth) != 2 || auth[0] != "Basic" {
			context.JSON(http.StatusUnauthorized,gin.H{
				"error": "Unauthorized",
			})
			context.Abort()
			return
		}

		//hacemos un decode de base 64
		credenctials, err := base64.StdEncoding.DecodeString(auth[1])
		if err != nil {
			log.Println(err)
		}

		//realizamos un Split por : y analizamos si el usuario o contraseña es igual
		// al que le pasamos a la funcion
		cr := strings.SplitN(string(credenctials),":",2)
		if len(cr) < 2 || strings.TrimSpace(cr[0]) != username || strings.TrimSpace(cr[1]) != password {
			context.JSON(http.StatusUnauthorized,gin.H{
				"error": "Unauthorized",
			})
			context.Abort()
			return
		}
		//si todo esta ok llegamos hasta el final y continuamos con la
		//ejecucion de la peticion
		context.Next()
	}
}
func main() {

	sumatoria := func(numbers ...int) []int{
		if numbers != nil {
			for index,_ := range numbers  {
				numbers[index] += 1
			}
		}
		return numbers
	}
	fmt.Println(sumatoria(1,2,3,4))


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

	//path con autorizacion requerida
	//authorization :=  server.Group("/person",basicAuth(CREDENTIAL_USERNAME,CREDENTIAL_PASSWORD))
	//authorization.GET("/", GetPerson)

	//path con autorizacion requerida
	authorization :=  server.Group("/person",gin.BasicAuth(gin.Accounts{"lagash":"meli",}))
	authorization.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{"messages":"Authorization Ok"})
	})


	// Inicio el servidor en el puerto 8080
	server.Run(":8080")
}