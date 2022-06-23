package domain

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

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error)
	LastID() (int64, error)
	Update(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error)
	UpdateSobrenomeIdade(id int64, sobrenome string, idade uint) (Usuario, error)
	Delete(id int64) error
}

type Service interface {
	GetAll() ([]Usuario, error)
	Store(nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error)
	Update(id int64, nome, sobrenome, email string, idade uint, altura float64, ativo bool, data string) (Usuario, error)
	UpdateSobrenomeIdade(id int64, sobrenome string, idade uint) (Usuario, error)
	Delete(id int64) error
}