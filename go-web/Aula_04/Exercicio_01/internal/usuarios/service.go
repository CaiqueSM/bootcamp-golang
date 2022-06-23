package usuarios

import (
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios/domain"
)

type service struct {
	repository domain.Repository
}

func NewService(r domain.Repository) domain.Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll() ([]domain.Usuario, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s service) Store(nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (domain.Usuario, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Usuario{}, err
	}

	lastID++

	usuario, err := s.repository.Store(lastID, nome, sobrenome, email, idade, altura, ativo, data)

	if err != nil {
		return domain.Usuario{}, err
	}

	return usuario, nil
}

func (s service) Update(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (domain.Usuario, error) {
	usuario, err := s.repository.Update(id, nome, sobrenome, email, idade, altura, ativo, data)
	if err != nil {
		return domain.Usuario{}, err
	}
	return usuario, err
}

func (s service) UpdateSobrenomeIdade(id int64, sobrenome string, idade uint) (domain.Usuario, error) {
	return s.repository.UpdateSobrenomeIdade(id, sobrenome, idade)
}

func (s service) Delete(id int64) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return err
}
