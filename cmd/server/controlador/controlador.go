package controlador

import (
	"os"
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
		t := os.Getenv("TOKEN")
		if token != t {
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

		p, errPost := c.service.Post(req)
		if errPost != nil {
			ctx.JSON(404, gin.H{
				"error": errPost.Error(),
			})
			return
		}

		ctx.JSON(200, p)
	}
}

func (c *Product) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//VALIDO TOKEN
		token := ctx.GetHeader("token")
		t := os.Getenv("TOKEN")
		if token != t {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		//VALIDO ID DEL HEADER
		id, idError := strconv.Atoi(ctx.Param("id"))
		if idError != nil {
			ctx.JSON(400, gin.H{
				"error": idError.Error(),
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
		// HAGO EL UPDATE
		p, errRes := c.service.Put(id, req)
		//SI FALLA ACA ES POR NOT FOUND
		if errRes != nil {
			ctx.JSON(404, gin.H{
				"error": errRes.Error(),
			})
			return
		}

		ctx.JSON(200, p)
	}
}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//VALIDO ID DEL HEADER
		id, idError := strconv.Atoi(ctx.Param("id"))
		if idError != nil {
			ctx.JSON(400, gin.H{
				"error": idError.Error(),
			})
			return
		}
		err := c.service.Delete(id)
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, nil)
	}
}

func (c *Product) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//VALIDO TOKEN
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		//VALIDO ID DEL HEADER
		id, idError := strconv.Atoi(ctx.Param("id"))
		if idError != nil {
			ctx.JSON(400, gin.H{
				"error": idError.Error(),
			})
			return
		}
		//VALIDO DATOS
		var req domain.ProductPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, gin.H{"error": "El precio del producto es requerido y debe ser mayor a cero"})
			return
		}
		// HAGO EL UPDATE
		p, errRes := c.service.Patch(id, req)
		//SI FALLA ACA ES POR NOT FOUND
		if errRes != nil {
			ctx.JSON(404, gin.H{
				"error": errRes.Error(),
			})
			return
		}

		ctx.JSON(200, p)
	}
}
