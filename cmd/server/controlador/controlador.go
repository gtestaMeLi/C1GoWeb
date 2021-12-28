package controlador

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gtestaMeLi/C1GoWeb/internal/domain"
	"github.com/gtestaMeLi/C1GoWeb/internal/products"
	"github.com/gtestaMeLi/C1GoWeb/pkg/web"
)

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Success 200 {object} web.Response
// @Router /Productos [get]
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		p := c.service.GetAll()
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// Products godoc
// @Summary Get product by id
// @Tags Products
// @Description get product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} web.Response
// @Router /Productos/{id} [get]
func (c *Product) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		p := c.service.Get(id)
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body domain.Product true "Product to store"
// @Success 200 {object} web.Response
// @Router /Productos [post]
func (c *Product) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//VALIDO TOKEN
		token := ctx.GetHeader("token")
		t := os.Getenv("TOKEN")
		if token != t {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		//VALIDO DATOS
		var req domain.Product
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		p, errPost := c.service.Post(req)
		if errPost != nil {
			ctx.JSON(404, web.NewResponse(404, nil, errPost.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// ModifyProducts godoc
// @Summary Modify products
// @Tags Products
// @Description modify products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param token header string true "token"
// @Param product body domain.Product true "Product to modify"
// @Success 200 {object} web.Response
// @Router /Productos/{id} [put]
func (c *Product) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//VALIDO TOKEN
		token := ctx.GetHeader("token")
		t := os.Getenv("TOKEN")
		if token != t {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		//VALIDO ID DEL HEADER
		id, idError := strconv.Atoi(ctx.Param("id"))
		if idError != nil {
			ctx.JSON(400, web.NewResponse(400, nil, idError.Error()))
			return
		}
		//VALIDO DATOS
		var req domain.Product
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		// HAGO EL UPDATE
		p, errRes := c.service.Put(id, req)
		//SI FALLA ACA ES POR NOT FOUND
		if errRes != nil {
			ctx.JSON(404, web.NewResponse(404, nil, errRes.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// DeleteProducts godoc
// @Summary Delete products
// @Tags Products
// @Description delete products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /Productos/{id} [delete]
func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//VALIDO ID DEL HEADER
		id, idError := strconv.Atoi(ctx.Param("id"))
		if idError != nil {
			ctx.JSON(400, web.NewResponse(400, nil, idError.Error()))
			return
		}
		err := c.service.Delete(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, nil, ""))
	}
}

// PatchProducts godoc
// @Summary Patch products
// @Tags Products
// @Description patch products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param token header string true "token"
// @Param product body domain.Product true "ProductPatch to patch"
// @Success 200 {object} web.Response
// @Router /Productos/{id} [patch]
func (c *Product) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//VALIDO TOKEN
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		//VALIDO ID DEL HEADER
		id, idError := strconv.Atoi(ctx.Param("id"))
		if idError != nil {
			ctx.JSON(400, web.NewResponse(400, nil, idError.Error()))
			return
		}
		//VALIDO DATOS
		var req domain.ProductPatch
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido y debe ser mayor a cero"))
			return
		}
		// HAGO EL UPDATE
		p, errRes := c.service.Patch(id, req)
		//SI FALLA ACA ES POR NOT FOUND
		if errRes != nil {
			ctx.JSON(404, web.NewResponse(404, nil, errRes.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}
