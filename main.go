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
	Nombre        string  `json:"nombre" binding:"required"`
	Apellido      string  `json:"apellido" binding:"required"`
	Email         string  `json:"email" binding:"required"`
	Edad          int64   `json:"edad" binding:"required"`
	Altura        float32 `json:"altura" binding:"required"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fecha-creacion" binding:"required"`
}

func main() {

	//OBTENEMOS LOS DATOS DEL JSON A MODO DE TENER DATOS PERSISTIDOS
	usuarios := []Usuarios{}
	jsonData, _ := os.ReadFile("./usuarios.json")
	if err := json.Unmarshal([]byte(jsonData), &usuarios); err != nil {
		log.Fatal(err)
	}

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

	//METODO POST
	router.POST("/Productos", func(c *gin.Context) {
		//VALIDO TOKEN
		token := c.GetHeader("token")
		if token != "111222" {
			c.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		//VALIDO DATOS
		var req Usuarios
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		req.ID = usuarios[len(usuarios)-1].ID + 1
		usuarios = append(usuarios, req)
		c.JSON(200, req)
	})

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
