package usuarios

import (
	"fmt"
	"log"

	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/pkg/store"
)

type Usuario struct {
	Id        int64   `json:"id"`
	Nome      string  `json:"nome"`
	Sobrenome string  `json:"sobrenome"`
	Email     string  `json:"email"`
	Idade     uint    `json:"idade"`
	Altura    float64 `json:"altura"`
	Ativo     bool    `json:"ativo"`
	Data      string  `json:"data"`
}

var ps []Usuario = []Usuario{}

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error)
	LastID() (int64, error)
	Update(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error)
	UpdateSobrenomeIdade(id int64, sobrenome string, idade uint) (Usuario, error)
	Delete(id int64) error
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Usuario, error) {
	var ps []Usuario
	r.db.Read(&ps)
	return ps, nil
}

func (r *repository) Store(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error) {
	var ps []Usuario
	r.db.Read(&ps)
	u := Usuario{id, nome, sobrenome, email, idade, altura, ativo, data}
	ps = append(ps, u)
	if err := r.db.Write(ps); err != nil {
		return Usuario{}, err
	}
	return u, nil
}

func (r *repository) LastID() (int64, error) {
	var ps []Usuario
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].Id, nil
}

func (repository) Update(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error) {
	u := Usuario{Nome: nome, Sobrenome: sobrenome, Email: email, Idade: idade, Altura: altura, Ativo: ativo, Data: data}
	updated := false
	log.Printf("%d", len(ps))
	for i := range ps {
		if ps[i].Id == id {
			u.Id = id
			ps[i] = u
			updated = true
		}
	}

	if !updated {
		return Usuario{}, fmt.Errorf("usuário %d não encontrado yo", id)
	}
	return u, nil
}

func (repository) UpdateSobrenomeIdade(id int64, sobrenome string, idade uint) (Usuario, error) {
	var u Usuario
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Sobrenome = sobrenome
			ps[i].Idade = idade
			updated = true
			u = ps[i]
		}
	}

	if !updated {
		return Usuario{}, fmt.Errorf("usuário %d não encontrado", id)
	}
	return u, nil
}

func (repository) Delete(id int64) error {
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
	return nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}