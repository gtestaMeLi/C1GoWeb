package productos

type Service interface {
	GetAll() []Product
	Get(id int) Product
	Post(prod Product) Product
}
type service struct {
	repo Repository
}

func NewService(s Repository) Service {
	return &service{
		repo: s,
	}
}
func (s *service) GetAll() []Product {
	ps := s.repo.GetAll()
	return ps
}

func (s *service) Get(id int) Product {
	p := s.repo.Get(id)
	return p
}

func (s *service) Post(prod Product) Product {

	p := s.repo.Post(prod)

	return p
}
