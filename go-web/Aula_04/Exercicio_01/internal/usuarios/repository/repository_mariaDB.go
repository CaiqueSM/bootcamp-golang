package repository

import (
	"database/sql"
	"log"

	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios/domain"
)

type repositoryMariaDB struct {
	db *sql.DB
}

// Delete implements domain.Repository
func (r *repositoryMariaDB) Delete(id int64) error {
	stmt, err := r.db.Prepare(`DELETE usuarios WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	if _, err = stmt.Exec(id); err != nil {
		return err
	}
	return nil
}

// GetAll implements domain.Repository
func (r *repositoryMariaDB) GetAll() ([]domain.Usuario, error) {
	var usuarios []domain.Usuario
	var usuario domain.Usuario
	rows, err := r.db.Query(`SELECT id, nome, sobrenome, email, idade, altura, ativo, data
	FROM usuarios`)
	if err != nil {
		log.Println(err)
		return []domain.Usuario{}, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&usuario.Id,
			&usuario.Nome,
			&usuario.Sobrenome,
			&usuario.Email,
			&usuario.Id,
			&usuario.Altura,
			&usuario.Ativo,
			&usuario.Data,
		); err != nil {
			log.Println(err.Error())
			return []domain.Usuario{}, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// LastID implements domain.Repository
func (r *repositoryMariaDB) LastID() (int64, error) {
	lastID:= 0
	rows, err := r.db.Query(`SELECT MAX(id) FROM usuarios`)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer rows.Close()
	if rows.Next(){
		if err := rows.Scan(&lastID); err != nil {
			log.Println(err.Error())
			return 0, err
		}
	}
	return int64(lastID), nil
}

// UpdateSobrenomeIdade implements domain.Repository
func (r *repositoryMariaDB) UpdateSobrenomeIdade(id int64, sobrenome string, idade uint) (domain.Usuario, error) {
	stmt, err := r.db.Prepare(`UPDATE usuarios SET, sobrenome = ?, idade = ? WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(sobrenome, idade, id)
	if err != nil {
		return domain.Usuario{}, err
	}
	return domain.Usuario{Id: id, Sobrenome: sobrenome, Idade: idade}, nil
}

func (r *repositoryMariaDB) Store(id int64,
	nome, sobrenome, email string,
	idade uint,
	altura float64,
	ativo bool,
	data string) (domain.Usuario, error) {
	stmt, err := r.db.Prepare(`INSERT INTO usuarios(nome, sobrenome, email, idade, altura, ativo, data)
		VALUES(?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(nome, sobrenome, email, idade, altura, ativo, data)
	if err != nil {
		return domain.Usuario{}, nil
	}
	log.Println(result.RowsAffected())
	return domain.Usuario{Id: id, Nome: nome, Sobrenome: sobrenome, Email: email, Idade: idade, Altura: altura, Ativo: ativo, Data: data}, nil

}

func (r *repositoryMariaDB) GetOne(id int) domain.Usuario {
	var usuario domain.Usuario
	rows, err := r.db.Query(`SELECT id, nome, sobrenome, email, idade, altura, ativo, data
		FROM usuarios
		WHERE id = ?`, id)
	if err != nil {
		log.Println(err)
		return usuario
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&usuario.Id,
			&usuario.Nome,
			&usuario.Sobrenome,
			&usuario.Email,
			&usuario.Id,
			&usuario.Altura,
			&usuario.Ativo,
			&usuario.Data,
		); err != nil {
			log.Println(err.Error())
			return usuario
		}
	}
	return usuario
}

func (r *repositoryMariaDB) Update(id int64,
	nome, sobrenome, email string,
	idade uint,
	altura float64,
	ativo bool,
	data string) (domain.Usuario, error) {
	stmt, err := r.db.Prepare(`UPDATE usuarios SET nome = ?, sobrenome = ?, email = ?, idade = ?, altura = ?, ativo = ?, data = ?
		WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(nome, sobrenome, email, idade, altura, ativo, data, id)
	if err != nil {
		return domain.Usuario{}, err
	}
	return domain.Usuario{Id: id, Nome: nome, Sobrenome: sobrenome, Email: email, Idade: idade, Altura: altura, Ativo: ativo, Data: data}, nil
}

func NewRepositoryMariaDB(db *sql.DB) domain.Repository {
	return &repositoryMariaDB{
		db: db,
	}
}
