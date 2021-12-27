package products

import "github.com/gtestaMeLi/C1GoWeb/internal/domain"

type Service interface {
	GetAll() []domain.Product
	Get(id int) domain.Product
	Post(prod domain.Product) domain.Product
	Put(id int, prod domain.Product) (domain.Product, error)
}
type service struct {
	repo Repository
}

func NewService(s Repository) Service {
	return &service{
		repo: s,
	}
}
func (s *service) GetAll() []domain.Product {
	ps := s.repo.GetAll()
	return ps
}

func (s *service) Get(id int) domain.Product {
	p := s.repo.Get(id)
	return p
}

func (s *service) Post(prod domain.Product) domain.Product {

	p := s.repo.Post(prod)

	return p
}

func (s *service) Put(id int, prod domain.Product) (res domain.Product, err error) {
	p, err := s.repo.Put(id, prod)

	return p, err
}
