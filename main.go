package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Usuarios struct {
	ID            int64   `json:"id"`
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
		c.JSON(200, usuarios)
	})

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
