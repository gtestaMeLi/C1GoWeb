package products

import (
	"fmt"

	"github.com/gtestaMeLi/C1GoWeb/internal/domain"
	"github.com/gtestaMeLi/C1GoWeb/pkg/store"
)

var p1 domain.Product = domain.Product{1, "Macbook", "pc", 100, 25000}
var p2 domain.Product = domain.Product{2, "Teclado", "accesorios pc", 200, 150}
var p3 domain.Product = domain.Product{3, "Monitor", "monitor", 10, 5000}

//var productos []domain.Product = []domain.Product{p1, p2, p3}

type Repository interface {
	GetAll() []domain.Product
	Get(id int) domain.Product
	Post(prod domain.Product) (domain.Product, error)
	Put(id int, prod domain.Product) (domain.Product, error)
	Delete(id int) error
	Patch(id int, p domain.ProductPatch) (domain.Product, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() []domain.Product {
	var productos []domain.Product

	r.db.Read(&productos)

	return productos
}

func (r *repository) Get(id int) domain.Product {
	resultado := domain.Product{}
	//extraigo los datos del archivo
	var productos []domain.Product
	r.db.Read(&productos)
	//lo busco
	for _, value := range productos {
		if id == value.ID {
			resultado = value
		}
	}

	return resultado
}

func (r *repository) Post(prod domain.Product) (domain.Product, error) {
	//extraigo los datos del archivo
	var productos []domain.Product
	r.db.Read(&productos)
	prod.ID = productos[len(productos)-1].ID + 1
	//productos = append(productos, prod)
	productos = append(productos, prod)
	if err := r.db.Write(productos); err != nil {
		return domain.Product{}, err
	}

	return prod, nil
}

func removeFromSlice(slice []domain.Product, s int) []domain.Product {
	return append(slice[:s], slice[s+1:]...)
}

func (r *repository) Put(id int, prod domain.Product) (res domain.Product, err error) {
	//extraigo los datos del archivo
	var productos []domain.Product
	r.db.Read(&productos)
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
	} else {
		if err := r.db.Write(productos); err != nil {
			return domain.Product{}, err
		}
	}
	return prod, nil
}

func (r *repository) Delete(id int) error {
	//extraigo los datos del archivo
	var productos []domain.Product
	var result []domain.Product
	r.db.Read(&productos)
	copy(result, productos)
	founded := false
	for i := range productos {
		if productos[i].ID == id {
			founded = true
			result = removeFromSlice(productos, i)
		}
	}
	if !founded {
		return fmt.Errorf("Producto %d no encontrado", id)
	} else {
		if err := r.db.Write(result); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) Patch(id int, p domain.ProductPatch) (domain.Product, error) {
	//extraigo los datos del archivo
	var productos []domain.Product
	r.db.Read(&productos)
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
	} else {
		if err := r.db.Write(productos); err != nil {
			return domain.Product{}, err
		}
	}
	return result, nil
}
