package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gtestaMeLi/C1GoWeb/cmd/server/controlador"
	"github.com/gtestaMeLi/C1GoWeb/docs"
	"github.com/gtestaMeLi/C1GoWeb/internal/products"
	"github.com/gtestaMeLi/C1GoWeb/pkg/store"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	_ = godotenv.Load()
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := controlador.NewProduct(service)
	r := gin.Default()

	//Documentacion de swagger
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
