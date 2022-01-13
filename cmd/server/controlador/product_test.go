package controlador

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gtestaMeLi/C1GoWeb/internal/domain"
	"github.com/gtestaMeLi/C1GoWeb/internal/products"
	"github.com/gtestaMeLi/C1GoWeb/pkg/store"
	"github.com/stretchr/testify/assert"
)

type responseCollections struct {
	Code  string           `json:"code"`
	Data  []domain.Product `json:"data"`
	Error string           `json:"error"`
}

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.Get())
	pr.POST("/", p.Post())
	pr.PUT("/:id", p.Put())
	pr.DELETE("/:id", p.Delete())
	pr.PATCH("/:id", p.Patch())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetProduct_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	objRes := responseCollections{}

	assert.Equal(t, 200, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	fmt.Println(objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)
}

func Test_SaveProduct_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/products/", `{
        "nombre": "Tester","tipo": "Funcional","cantidad": 10,"precio": 99.99
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func Test_UpdateProduct_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo PUT y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPatch, "/products/3", `{
        "nombre": "Tester","precio": 99.99
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func Test_DeleteProduct_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo PUT y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodDelete, "/products/4", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 404, rr.Code)
}
