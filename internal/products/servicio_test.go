package products

import (
	"testing"

	"github.com/gtestaMeLi/C1GoWeb/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestServicePatch(t *testing.T) {
	producto1 := domain.Product{1, "P1", "type1", 190, 1900}
	producto2 := domain.Product{2, "P2", "type2", 190, 1900}
	aux := []domain.Product{producto1, producto2}
	db := mockStore{false, aux}
	repo := NewRepository(&db)
	service := NewService(repo)

	//Resultado esperado
	esperado := domain.Product{1, "After Product", "type1", 190, 1900}
	cambios := domain.ProductPatch{"After Product", 1900}

	///

	resultado, err := service.Patch(1, cambios)
	assert.Nil(t, err, "hubo un error en el update")
	assert.Equal(t, esperado, resultado, "deben ser iguales")
	assert.Equal(t, true, db.readExecute, "no se ejecuto el read")
}

func TestServiceDelete(t *testing.T) {
	producto1 := domain.Product{1, "P1", "type1", 190, 1900}
	producto2 := domain.Product{2, "P2", "type2", 190, 1900}
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
