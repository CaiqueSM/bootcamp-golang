package usuarios

import "fmt"

type Usuario struct {
	Id        int   `json:"id"`
	Nome      string  `json:"nome" binding:"required"`
	Sobrenome string  `json:"sobrenome" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Idade     uint    `json:"idade" binding:"required"`
	Altura    float64 `json:"altura" binding:"required"`
	Ativo     bool    `json:"ativo" binding:"required"`
	Data      string  `json:"data" binding:"required"`
}

var ps []Usuario
var lastID int

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(id int, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error)
	LastID() (int, error)
	Update(id int, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Usuario, error) {
	return ps, nil
}

func (r *repository) Store(id int, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error) {
	u := Usuario{id, nome, sobrenome, email, idade, altura, ativo, data}
	ps = append(ps, u)
	lastID = u.Id
	return u, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Update(id int, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error){
	u:= Usuario{id, nome, sobrenome, email, idade, altura, ativo, data}
	updated := false
	for i:= range ps{
		if ps[i].Id == id{
			u.Id = id
			ps[i] = u
			updated = true
		}
	}

	if !updated{
		return Usuario{}, fmt.Errorf("usuário %d não encontrado", id)
	}
	return u, nil
}

