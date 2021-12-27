package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gtestaMeLi/C1GoWeb/cmd/server/controlador"
	"github.com/gtestaMeLi/C1GoWeb/internal/products"
)

func main() {

	repo := products.NewRepository()
	service := products.NewService(repo)
	p := controlador.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/Productos")
	{
		pr.GET("/", p.GetAll())
		pr.GET("/:id", p.Get())
		pr.POST("/", p.Post())
		pr.PUT("/:id", p.Put())
		pr.DELETE("/:id", p.Delete())
		pr.PATCH("/:id", p.Patch())
	}

	r.Run()
}
