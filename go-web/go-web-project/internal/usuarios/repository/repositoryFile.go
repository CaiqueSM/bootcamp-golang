package repository

import (
	"fmt"

	"github.com/CaiqueSM/bootcamp-golang.git/go-web/go-web-project/pkg/store"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/go-web-project/internal/usuarios/domain"
)

type repositoryFile struct {
	db store.Store
}

func NewRepository(db store.Store) domain.Repository {
	return &repositoryFile{
		db: db,
	}
}

func (r *repositoryFile) GetAll() ([]domain.Usuario, error) {
	var ps []domain.Usuario
	if err := r.db.Read(&ps); err != nil {
		return []domain.Usuario{}, err
	}
	return ps, nil
}

func (r *repositoryFile) Store(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (domain.Usuario, error) {
	var ps []domain.Usuario
	if err := r.db.Read(&ps); err != nil {
		return domain.Usuario{}, err
	}
	u := domain.Usuario{Id: id, Nome: nome, Sobrenome: sobrenome, Email: email, Idade: idade, Altura: altura, Ativo: ativo, Data: data}
	ps = append(ps, u)
	if err := r.db.Write(ps); err != nil {
		return domain.Usuario{}, err
	}
	return u, nil
}

func (r *repositoryFile) LastID() (int64, error) {
	var ps []domain.Usuario
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].Id, nil
}

func (r *repositoryFile) Update(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (domain.Usuario, error) {
	u := domain.Usuario{Nome: nome, Sobrenome: sobrenome, Email: email, Idade: idade, Altura: altura, Ativo: ativo, Data: data}
	var ps []domain.Usuario
	if err := r.db.Read(&ps); err != nil {
		return domain.Usuario{}, err
	}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			u.Id = id
			ps[i] = u
			updated = true
		}
	}

	if !updated {
		return domain.Usuario{}, fmt.Errorf("usuário %d não encontrado", id)
	}

	if err := r.db.Write(ps); err != nil {
		return domain.Usuario{}, err
	}

	return u, nil
}

func (r *repositoryFile) UpdateSobrenomeIdade(id int64, sobrenome string, idade uint) (domain.Usuario, error) {
	var u domain.Usuario
	var ps []domain.Usuario
	if err := r.db.Read(&ps); err != nil {
		return domain.Usuario{}, err
	}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			if sobrenome != "" {
				ps[i].Sobrenome = sobrenome
			}
			if idade != 0 {
				ps[i].Idade = idade
			}

			updated = true
			u = ps[i]
		}
	}

	if !updated {
		return domain.Usuario{}, fmt.Errorf("usuário %d não encontrado", id)
	}

	if err := r.db.Write(ps); err != nil {
		return domain.Usuario{}, err
	}

	return u, nil
}

func (r *repositoryFile) Delete(id int64) error {
	var ps []domain.Usuario
	if err := r.db.Read(&ps); err != nil {
		return err
	}

	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("usuário %d não encontrado", id)
	}
	ps = append(ps[:index], ps[index+1:]...)

	if err := r.db.Write(ps); err != nil {
		return err
	}
	return nil
}
