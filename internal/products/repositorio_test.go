package products

import (
	"encoding/json"
	"testing"

	"github.com/gtestaMeLi/C1GoWeb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type stubStore struct{}

type mockStore struct {
	readExecute bool
	productList []domain.Product
}

func (s *stubStore) Read(data interface{}) error {
	producto1 := domain.Product{ID: 1, Name: "P1", Type: "type1", Count: 190, Price: 1900}
	producto2 := domain.Product{ID: 2, Name: "P2", Type: "type2", Count: 190, Price: 1900}

	aux, _ := json.Marshal([]domain.Product{producto1, producto2})

	err := json.Unmarshal(aux, &data)

	if err != nil {
		return err
	}

	return nil
}

func (s *mockStore) Read(data interface{}) error {

	aux, _ := json.Marshal(s.productList)

	err := json.Unmarshal(aux, &data)
	if err != nil {
		return err
	}
	s.readExecute = true
	return nil
}

func (s *stubStore) Write(data interface{}) error { return nil }
func (s *mockStore) Write(data interface{}) error {
	s.productList = nil
	fileData, _ := json.Marshal(data)
	err := json.Unmarshal(fileData, &s.productList)
	if err != nil {
		return err
	}
	return nil

}

func TestRead(t *testing.T) {
	db := stubStore{}
	repo := NewRepository(&db)

	//Resultado esperado
	producto1 := domain.Product{ID: 1, Name: "P1", Type: "type1", Count: 190, Price: 1900}
	producto2 := domain.Product{ID: 2, Name: "P2", Type: "type2", Count: 190, Price: 1900}

	esperado := []domain.Product{producto1, producto2}

	///

	resultado := repo.GetAll()
	assert.Equal(t, esperado, resultado, "deben ser iguales")

}

func TestGet(t *testing.T) {
	producto1 := domain.Product{ID: 1, Name: "P1", Type: "type1", Count: 190, Price: 1900}
	producto2 := domain.Product{ID: 2, Name: "P2", Type: "type2", Count: 190, Price: 1900}
	aux := []domain.Product{producto1, producto2}
	db := mockStore{false, aux}
	repo := NewRepository(&db)

	resultado := repo.Get(1)
	assert.Equal(t, producto1, resultado, "Deen ser iguales")
}

func TestPost(t *testing.T) {
	producto1 := domain.Product{ID: 1, Name: "P1", Type: "type1", Count: 190, Price: 1900}
	producto2 := domain.Product{ID: 2, Name: "P2", Type: "type2", Count: 190, Price: 1900}
	aux := []domain.Product{producto1}
	db := mockStore{false, aux}
	repo := NewRepository(&db)
	//Resultado esperado
	esperado := []domain.Product{producto1, producto2}

	///

	resultado, err := repo.Post(producto2)
	assert.Nil(t, err, "hubo un error en el post")
	assert.Equal(t, resultado, producto2, "Deben ser iguales")
	assert.Equal(t, esperado, db.productList, "deben ser iguales")
}

func TestPatch(t *testing.T) {
	producto1 := domain.Product{ID: 1, Name: "P1", Type: "type1", Count: 190, Price: 1900}
	producto2 := domain.Product{ID: 2, Name: "P2", Type: "type2", Count: 190, Price: 1900}
	aux := []domain.Product{producto1, producto2}
	db := mockStore{false, aux}
	repo := NewRepository(&db)

	//Resultado esperado
	esperado := domain.Product{ID: 1, Name: "After Product", Type: "type1", Count: 190, Price: 1900}
	cambios := domain.ProductPatch{Name: "After Product", Price: 1900}

	///

	resultado, err := repo.Patch(1, cambios)
	assert.Nil(t, err, "hubo un error en el update")
	assert.Equal(t, esperado, resultado, "deben ser iguales")
	assert.Equal(t, true, db.readExecute, "no se ejecuto el read")
}

func TestPut(t *testing.T) {
	producto1 := domain.Product{ID: 1, Name: "P1", Type: "type1", Count: 190, Price: 1900}
	producto2 := domain.Product{ID: 2, Name: "P2", Type: "type2", Count: 190, Price: 1900}
	aux := []domain.Product{producto1, producto2}
	db := mockStore{false, aux}
	repo := NewRepository(&db)

	//Resultado esperado
	esperado := domain.Product{ID: 1, Name: "After Product", Type: "type1", Count: 190, Price: 1900}

	///

	resultado, err := repo.Put(1, esperado)
	assert.Nil(t, err, "hubo un error en el update")
	assert.Equal(t, esperado, resultado, "deben ser iguales")
	_, err2 := repo.Put(8, producto1)
	assert.NotNil(t, err2, "hubo un error en el update")

}
