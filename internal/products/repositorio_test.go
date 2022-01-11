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
}

func (s *stubStore) Read(data interface{}) error {
	producto1 := domain.Product{1, "P1", "type1", 190, 1900}
	producto2 := domain.Product{2, "P2", "type2", 190, 1900}

	aux, _ := json.Marshal([]domain.Product{producto1, producto2})

	json.Unmarshal(aux, &data)

	return nil
}

func (s *mockStore) Read(data interface{}) error {
	producto1 := domain.Product{1, "Before Product", "type1", 190, 1900}
	producto2 := domain.Product{2, "P2", "type2", 190, 1900}

	aux, _ := json.Marshal([]domain.Product{producto1, producto2})

	json.Unmarshal(aux, &data)
	s.readExecute = true
	return nil
}

func (s *stubStore) Write(data interface{}) error { return nil }
func (s *mockStore) Write(data interface{}) error { return nil }

// func TestRead(t *testing.T) {
// 	db := stubStore{}
// 	repo := NewRepository(&db)

// 	//Resultado esperado
// 	producto1 := domain.Product{1, "P1", "type1", 190, 1900}
// 	producto2 := domain.Product{2, "P2", "type2", 190, 1900}

// 	esperado := []domain.Product{producto1, producto2}

// 	///

// 	resultado := repo.GetAll()
// 	assert.Equal(t, esperado, resultado, "deben ser iguales")

// }

func TestPatch(t *testing.T) {
	db := mockStore{}
	repo := NewRepository(&db)

	//Resultado esperado
	esperado := domain.Product{1, "After Product", "type1", 190, 1900}
	cambios := domain.ProductPatch{"After Product", 1900}

	///

	resultado, err := repo.Patch(1, cambios)
	assert.Nil(t, err, "hubo un error en el update")
	assert.Equal(t, esperado, resultado, "deben ser iguales")
	assert.Equal(t, true, db.readExecute, "no se ejecuto el read")
}
