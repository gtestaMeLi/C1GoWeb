package controlador

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gtestaMeLi/C1GoWeb/internal/domain"
	"github.com/gtestaMeLi/C1GoWeb/internal/products"
)

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		p := c.service.GetAll()
		ctx.JSON(200, p)
	}
}
func (c *Product) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		p := c.service.Get(id)
		ctx.JSON(200, p)
	}
}

func (c *Product) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//VALIDO TOKEN
		token := ctx.GetHeader("token")
		if token != "111222" {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		//VALIDO DATOS
		var req domain.Product
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		p := c.service.Post(req)
		ctx.JSON(200, p)
	}
}
