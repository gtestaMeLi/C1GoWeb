package products

import (
	"testing"

	"github.com/gtestaMeLi/C1GoWeb/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestServicePatch(t *testing.T) {
	producto1 := domain.Product{ID: 1, Name: "P1", Type: "type1", Count: 190, Price: 1900}
	producto2 := domain.Product{ID: 2, Name: "P2", Type: "type2", Count: 190, Price: 1900}
	aux := []domain.Product{producto1, producto2}
	db := mockStore{false, aux}
	repo := NewRepository(&db)
	service := NewService(repo)

	//Resultado esperado
	esperado := domain.Product{ID: 1, Name: "After Product", Type: "type1", Count: 190, Price: 1900}
	cambios := domain.ProductPatch{Name: "After Product", Price: 1900}

	///

	resultado, err := service.Patch(1, cambios)
	assert.Nil(t, err, "hubo un error en el update")
	assert.Equal(t, esperado, resultado, "deben ser iguales")
	assert.Equal(t, true, db.readExecute, "no se ejecuto el read")
}

func TestServiceGetAll(t *testing.T) {
	db := stubStore{}
	repo := NewRepository(&db)
	service := NewService(repo)

	//Resultado esperado
	producto1 := domain.Product{ID: 1, Name: "P1", Type: "type1", Count: 190, Price: 1900}
	producto2 := domain.Product{ID: 2, Name: "P2", Type: "type2", Count: 190, Price: 1900}

	esperado := []domain.Product{producto1, producto2}

	///

	resultado := service.GetAll()
	assert.Equal(t, esperado, resultado, "deben ser iguales")

}

func TestServiceGet(t *testing.T) {
	producto1 := domain.Product{ID: 1, Name: "P1", Type: "type1", Count: 190, Price: 1900}
	producto2 := domain.Product{ID: 2, Name: "P2", Type: "type2", Count: 190, Price: 1900}
	aux := []domain.Product{producto1, producto2}
	db := mockStore{false, aux}
	repo := NewRepository(&db)
	service := NewService(repo)

	resultado := service.Get(1)
	assert.Equal(t, producto1, resultado, "Deen ser iguales")
}

func TestServicePut(t *testing.T) {
	producto1 := domain.Product{ID: 1, Name: "P1", Type: "type1", Count: 190, Price: 1900}
	producto2 := domain.Product{ID: 2, Name: "P2", Type: "type2", Count: 190, Price: 1900}
	aux := []domain.Product{producto1, producto2}
	db := mockStore{false, aux}
	repo := NewRepository(&db)
	service := NewService(repo)

	//Resultado esperado
	esperado := domain.Product{ID: 1, Name: "After Product", Type: "type1", Count: 190, Price: 1900}

	///

	resultado, err := service.Put(1, esperado)
	assert.Nil(t, err, "hubo un error en el update")
	assert.Equal(t, esperado, resultado, "deben ser iguales")
	_, err2 := service.Put(8, producto1)
	assert.NotNil(t, err2, "hubo un error en el update")

}

func TestServiceDelete(t *testing.T) {
	producto1 := domain.Product{ID: 1, Name: "P1", Type: "type1", Count: 190, Price: 1900}
	producto2 := domain.Product{ID: 2, Name: "P2", Type: "type2", Count: 190, Price: 1900}
	aux := []domain.Product{producto1, producto2}
	db := mockStore{false, aux}
	repo := NewRepository(&db)
	service := NewService(repo)

	///Lista esperada al eliminar un valor
	result := []domain.Product{producto2}

	existProductError := service.Delete(1)
	assert.Nil(t, existProductError, "hubo un error en el delete")
	assert.Equal(t, result, db.productList, "no se elimino correctamente el producto")
	notExistsProductError := service.Delete(5)
	assert.NotNil(t, notExistsProductError, "deben ser iguales")

}

func TestServicePost(t *testing.T) {
	producto1 := domain.Product{ID: 1, Name: "P1", Type: "type1", Count: 190, Price: 1900}
	producto2 := domain.Product{ID: 2, Name: "P2", Type: "type2", Count: 190, Price: 1900}
	aux := []domain.Product{producto1}
	db := mockStore{false, aux}
	repo := NewRepository(&db)
	service := NewService(repo)
	//Resultado esperado
	esperado := []domain.Product{producto1, producto2}

	///

	resultado, err := service.Post(producto2)
	assert.Nil(t, err, "hubo un error en el post")
	assert.Equal(t, resultado, producto2, "Deben ser iguales")
	assert.Equal(t, esperado, db.productList, "deben ser iguales")
}
