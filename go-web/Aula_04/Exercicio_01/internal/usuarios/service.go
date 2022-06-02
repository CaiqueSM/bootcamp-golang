package usuarios

type Service interface {
	GetAll() ([]Usuario, error)
	Store(nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error)
	Update(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error)
	UpdateSobrenome(id int64, sobrenome string)(Usuario, error)
	UpdateIdade(id int64, idade uint)(Usuario, error)
	Delete(id int64) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Usuario, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Usuario{}, err
	}

	lastID++

	usuario, err := s.repository.Store(lastID, nome, sobrenome, email, idade, altura, ativo, data)

	if err != nil {
		return Usuario{}, err
	}

	return usuario, nil
}

func (s *service)Update(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error){
	return s.repository.Update(id, nome, sobrenome, email, idade, altura, ativo, data)
}

func (s *service)UpdateSobrenome(id int64, sobrenome string)(Usuario, error){
	return s.repository.UpdateSobrenome(id, sobrenome)
}

func (s *service)UpdateIdade(id int64, idade uint)(Usuario, error){
	return s.repository.UpdateIdade(id, idade)
}

func (s *service) Delete(id int64) error{
	return s.repository.Delete(id)
}