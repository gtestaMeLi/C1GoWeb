package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Usuarios struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int64   `json:"edad"`
	Altura        float32 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fecha-creacion"`
}

func main() {

	// Crea un router con gin
	router := gin.Default()
	// // Captura la solicitud GET
	router.GET("/holanombre", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Gianni",
		})
	})

	// Captura la solicitud GETALL
	router.GET("/Productos/GetAll", func(c *gin.Context) {
		usuarios := []Usuarios{}
		jsonData, _ := os.ReadFile("./usuarios.json")
		if err := json.Unmarshal([]byte(jsonData), &usuarios); err != nil {
			log.Fatal(err)
		}
		resultado := []Usuarios{}
		//buscamos los query params en el contexto
		var filterName = c.Query("nombre")
		if filterName != "" {
			for _, value := range usuarios {
				if strings.Contains(value.Nombre, filterName) {
					resultado = append(resultado, value)
				}
			}
		} else {
			resultado = usuarios
		}

		c.JSON(200, resultado)
	})

	router.GET("/Productos/:id", func(c *gin.Context) {

		usuarios := []Usuarios{}
		jsonData, _ := os.ReadFile("./usuarios.json")
		if err := json.Unmarshal([]byte(jsonData), &usuarios); err != nil {
			log.Fatal(err)
		}
		resultado := Usuarios{}

		id, _ := strconv.Atoi(c.Param("id"))

		for _, value := range usuarios {
			if id == value.ID {
				resultado = value
			}
		}

		c.JSON(200, resultado)
	})

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
