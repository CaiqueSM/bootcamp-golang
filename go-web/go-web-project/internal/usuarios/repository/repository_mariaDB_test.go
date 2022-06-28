package repository_test

import (
	"testing"

	"github.com/CaiqueSM/bootcamp-golang.git/go-web/go-web-project/internal/usuarios/domain"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/go-web-project/internal/usuarios/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestStoreDB(t *testing.T) {
	usuario := domain.Usuario{
		Id:        1,
		Nome:      "ABF",
		Sobrenome: "DEFG",
		Email:     "generico@Email.com",
		Idade:     25,
		Altura:    1.68,
		Ativo:     true,
		Data:      "23-06-2022",
	}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(`INSERT INTO usuarios`)
	mock.ExpectExec(`INSERT INTO usuarios`).WillReturnResult(sqlmock.NewResult(1, 1))

	myRepo := repository.NewRepositoryMariaDB(db)
	usuarioResult, err := myRepo.Store(1, "ABF", "DEFG", "generico@Email.com", 25, 1.68, true, "23-06-2022")
	assert.NoError(t, err)
	assert.Equal(t, usuario, usuarioResult)
}

func TestGetAllDB(t *testing.T) {
	usuario := domain.Usuario{
		Id:        1,
		Nome:      "ABF",
		Sobrenome: "DEFG",
		Email:     "generico@Email.com",
		Idade:     25,
		Altura:    1.68,
		Ativo:     true,
		Data:      "23-06-2022",
	}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	usuarioList := []domain.Usuario{}
	usuarioList = append(usuarioList, usuario)

	columns := []string{"id", "nome", "sobrenome", "email", "idade", "altura", "ativo", "data"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(
		usuarioList[0].Id,
		usuarioList[0].Nome,
		usuarioList[0].Sobrenome,
		usuarioList[0].Email,
		usuarioList[0].Idade,
		usuarioList[0].Altura,
		usuarioList[0].Ativo,
		usuarioList[0].Data,
	)

	mock.ExpectQuery("SELECT .* FROM usuarios").WillReturnRows(rows)
	myRepo := repository.NewRepositoryMariaDB(db)
	usuarioResultList, err := myRepo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, usuarioList, usuarioResultList)
}

func TestUpdateDB(t *testing.T) {
	usuarioAtualizado := domain.Usuario{
		Id:        1,
		Nome:      "ABF",
		Sobrenome: "After Update",
		Email:     "generico@Email.com",
		Idade:     26,
		Altura:    1.68,
		Ativo:     true,
		Data:      "23-06-2022",
	}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("UPDATE usuarios")
	mock.ExpectExec("UPDATE usuarios").WillReturnResult(sqlmock.NewResult(1,1))

	myRepo := repository.NewRepositoryMariaDB(db)
	usuarioResult, err := myRepo.Update(1, "ABF", "After Update", "generico@Email.com", 26, 1.68, true, "23-06-2022")
	assert.NoError(t, err)
	assert.Equal(t, usuarioAtualizado, usuarioResult)
}
