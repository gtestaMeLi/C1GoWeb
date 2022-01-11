package products

import (
	"encoding/json"
	"testing"

	"github.com/gtestaMeLi/C1GoWeb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type stubStore struct{}

func (s *stubStore) Read(data interface{}) error {
	producto1 := domain.Product{1, "P1", "type1", 190, 1900}
	producto2 := domain.Product{2, "P2", "type2", 190, 1900}

	aux, _ := json.Marshal([]domain.Product{producto1, producto2})

	json.Unmarshal(aux, &data)
	return nil
}

func (s *stubStore) Write(data interface{}) error { return nil }

func TestRead(t *testing.T) {
	db := stubStore{}
	repo := NewRepository(&db)

	//Resultado esperado
	producto1 := domain.Product{1, "P1", "type1", 190, 1900}
	producto2 := domain.Product{2, "P2", "type2", 190, 1900}

	esperado := []domain.Product{producto1, producto2}

	///

	resultado := repo.GetAll()
	assert.Equal(t, esperado, resultado, "deben ser iguales")

}
