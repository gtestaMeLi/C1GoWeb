package products

import (
	"fmt"

	"github.com/gtestaMeLi/C1GoWeb/internal/domain"
)

var p1 domain.Product = domain.Product{1, "Macbook", "pc", 100, 25000}
var p2 domain.Product = domain.Product{2, "Teclado", "accesorios pc", 200, 150}
var p3 domain.Product = domain.Product{3, "Monitor", "monitor", 10, 5000}

var productos []domain.Product = []domain.Product{p1, p2, p3}

type Repository interface {
	GetAll() []domain.Product
	Get(id int) domain.Product
	Post(prod domain.Product) domain.Product
	Put(id int, prod domain.Product) (domain.Product, error)
	Delete(id int) error
	Patch(id int, p domain.ProductPatch) (domain.Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() []domain.Product {

	return productos
}

func (r *repository) Get(id int) domain.Product {
	resultado := domain.Product{}

	for _, value := range productos {
		if id == value.ID {
			resultado = value
		}
	}

	return resultado
}

func (r *repository) Post(prod domain.Product) domain.Product {
	prod.ID = productos[len(productos)-1].ID + 1
	productos = append(productos, prod)

	return prod
}

func removeFromSlice(slice []domain.Product, s int) []domain.Product {
	return append(slice[:s], slice[s+1:]...)
}

func (r *repository) Put(id int, prod domain.Product) (res domain.Product, err error) {

	updated := false
	for i := range productos {
		if productos[i].ID == id {
			prod.ID = id
			productos[i] = prod
			updated = true
		}
	}
	if !updated {
		return domain.Product{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return prod, nil
}

func (r *repository) Delete(id int) error {
	founded := false
	for i := range productos {
		if productos[i].ID == id {
			founded = true
			productos = removeFromSlice(productos, i)
		}
	}
	if !founded {
		return fmt.Errorf("Producto %d no encontrado", id)
	}
	return nil
}

func (r *repository) Patch(id int, p domain.ProductPatch) (domain.Product, error) {
	updated := false
	result := domain.Product{}
	for i := range productos {
		if productos[i].ID == id {
			result = productos[i]
			result.Name = p.Name
			result.Price = p.Price
			productos[i] = result
			updated = true
		}
	}
	if !updated {
		return domain.Product{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return result, nil
}
