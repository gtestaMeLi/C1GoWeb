package productos

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var p1 Product = Product{1, "Macbook", "pc", 100, 25000}
var p2 Product = Product{2, "Teclado", "accesorios pc", 200, 150}
var p3 Product = Product{3, "Monitor", "monitor", 10, 5000}

var productos []Product = []Product{p1, p2, p3}

type Repository interface {
	GetAll() []Product
	Get(id int) Product
	Post(prod Product) Product
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() []Product {

	return productos
}

func (r *repository) Get(id int) Product {
	resultado := Product{}

	for _, value := range productos {
		if id == value.ID {
			resultado = value
		}
	}

	return resultado
}

func (r *repository) Post(prod Product) Product {
	prod.ID = productos[len(productos)-1].ID + 1
	productos = append(productos, prod)

	return prod
}
