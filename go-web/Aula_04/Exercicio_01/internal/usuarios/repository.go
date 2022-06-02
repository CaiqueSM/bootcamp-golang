package usuarios

import "fmt"

type Usuario struct {
	Id        int64   `json:"id"`
	Nome      string  `json:"nome" binding:"required"`
	Sobrenome string  `json:"sobrenome" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Idade     uint    `json:"idade" binding:"required"`
	Altura    float64 `json:"altura" binding:"required"`
	Ativo     bool    `json:"ativo" binding:"required"`
	Data      string  `json:"data" binding:"required"`
}

var ps []Usuario
var lastID int64

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error)
	LastID() (int64, error)
	Update(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error)
	UpdateSobrenome(id int64, sobrenome string)(Usuario, error)
	UpdateIdade(id int64, idade uint)(Usuario, error)
	Delete(id int64) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Usuario, error) {
	return ps, nil
}

func (r *repository) Store(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error) {
	u := Usuario{id, nome, sobrenome, email, idade, altura, ativo, data}
	ps = append(ps, u)
	lastID = u.Id
	return u, nil
}

func (r *repository) LastID() (int64, error) {
	return lastID, nil
}

func (r *repository) Update(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error){
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

func (r *repository)UpdateSobrenome(id int64, sobrenome string)(Usuario, error){
	var u Usuario
	updated := false
	for i:= range ps{
		if ps[i].Id == id{
			ps[i].Sobrenome = sobrenome
			updated = true
			u = ps[i]
		}
	}

	if !updated{
		return Usuario{}, fmt.Errorf("usuário %d não encontrado", id)
	}
	return u, nil
}

func (r *repository)UpdateIdade(id int64, idade uint)(Usuario, error){
	var u Usuario
	updated := false
	for i:= range ps{
		if ps[i].Id == id{
			ps[i].Idade = idade
			updated = true
			u = ps[i]
		}
	}

	if !updated{
		return Usuario{}, fmt.Errorf("usuário %d não encontrado", id)
	}
	return u, nil
}

func (r *repository) Delete(id int64) error{
	deleted := false
	var index int
	for i := range ps{
		if ps[i].Id == id{
			index = i
			deleted = true
		}
	}
	if !deleted{
		return fmt.Errorf("usuário %d não encontrado", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	return nil
}