package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gtestaMeLi/C1GoWeb/cmd/server/controlador"
	"github.com/gtestaMeLi/C1GoWeb/internal/products"
	"github.com/gtestaMeLi/C1GoWeb/pkg/store"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
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
